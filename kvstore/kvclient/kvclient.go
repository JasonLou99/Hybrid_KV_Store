package main

import (
	"context"
	"flag"
	"fmt"
	"hybrid_kv_store/rpc/kvrpc"
	"hybrid_kv_store/util"
	"math/rand"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

/*
	kvclient有三种方式进行通信，http协议、原生tcp协议、grpc通信
	前两种面向wasm runtime调用，后一种是常规的kvs通信模式，用于性能对比
*/
type KVClient struct {
	kvservers        []string
	vectorclock      map[string]int32
	consistencyLevel int32
	kvsId            int // target node
}

// func MakeKVClient(kvservers []string) *KVClient {
// 	return &KVClient{
// 		kvservers:   kvservers,
// 		vectorclock: make(map[string]int32),
// 	}
// }

const (
	CAUSAL = iota
	BoundedStaleness
)

// Method of Send RPC of GetInCausal
func (kvc *KVClient) SendGetInCausal(address string, request *kvrpc.GetInCasualRequest) (*kvrpc.GetInCasualResponse, error) {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		util.EPrintf("err in SendGetInCausal: %v", err)
		return nil, err
	}
	defer conn.Close()
	client := kvrpc.NewKVClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	reply, err := client.GetInCasual(ctx, request)
	if err != nil {
		util.EPrintf("err in SendGetInCausal: %v", err)
		return nil, err
	}
	return reply, nil
}

// Method of Send RPC of PutInCausal
func (kvc *KVClient) SendPutInCausal(address string, request *kvrpc.PutInCasualRequest) (*kvrpc.PutInCasualResponse, error) {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		util.EPrintf("err in SendPutInCausal: %v", err)
		return nil, err
	}
	defer conn.Close()
	client := kvrpc.NewKVClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	reply, err := client.PutInCasual(ctx, request)
	if err != nil {
		util.EPrintf("err in SendPutInCausal: %v", err)
		return nil, err
	}
	return reply, nil
}

// Client Get Value, Read One Replica
func (kvc *KVClient) GetInCausal(key string) (string, bool) {
	request := &kvrpc.GetInCasualRequest{
		Key:         key,
		Vectorclock: kvc.vectorclock,
	}
	for {
		request.Timestamp = time.Now().UnixMilli()
		reply, err := kvc.SendGetInCausal(kvc.kvservers[kvc.kvsId], request)
		if err != nil {
			util.EPrintf("err in GetInCausal: %v", err)
			return "", false
		}
		if reply.Vectorclock != nil && reply.Success {
			kvc.vectorclock = reply.Vectorclock
			return reply.Value, reply.Success
		}
		// refresh the target node
		// util.DPrintf("GetInCausal Failed, refresh the target node: %v", kvc.kvservers[kvc.kvsId])
		fmt.Println("GetInCausal Failed, refresh the target node: ", kvc.kvservers[kvc.kvsId])
		kvc.kvsId = (kvc.kvsId + 1) % len(kvc.kvservers)
		atomic.AddInt32(&falseTime, 1)
	}
}

// Client Get Value, Read Quorum Replica
func (kvc *KVClient) GetInCausalWithQuorum(key string) (string, bool) {
	request := &kvrpc.GetInCasualRequest{
		Key:         key,
		Vectorclock: kvc.vectorclock,
	}
	for {
		request.Timestamp = time.Now().UnixMilli()
		LatestReply := &kvrpc.GetInCasualResponse{
			Vectorclock: util.MakeMap(kvc.kvservers),
			Value:       "",
			Success:     false,
		}
		// Get Value From All Nodes (Get All, Put One)
		for i := 0; i < len(kvc.kvservers); i++ {
			reply, err := kvc.SendGetInCausal(kvc.kvservers[i], request)
			if err != nil {
				util.EPrintf("err in GetInCausalWithQuorum: %v", err)
				return "", false
			}
			// if reply is newer, update the LatestReply
			if util.IsUpper(util.BecomeSyncMap(reply.Vectorclock), util.BecomeSyncMap(LatestReply.Vectorclock)) {
				LatestReply = reply
			}
		}
		if LatestReply.Vectorclock != nil && LatestReply.Success {
			kvc.vectorclock = LatestReply.Vectorclock
			return LatestReply.Value, LatestReply.Success
		}
		// refresh the target node
		util.DPrintf("GetInCausalWithQuorum Failed, refresh the target node: %v", kvc.kvservers[kvc.kvsId])
		kvc.kvsId = (kvc.kvsId + 1) % len(kvc.kvservers)
	}
}

// Client Put Value
func (kvc *KVClient) PutInCausal(key string, value string) bool {
	request := &kvrpc.PutInCasualRequest{
		Key:         key,
		Value:       value,
		Vectorclock: kvc.vectorclock,
		Timestamp:   time.Now().UnixMilli(),
	}
	// keep sending PutInCausal until success
	for {
		reply, err := kvc.SendPutInCausal(kvc.kvservers[kvc.kvsId], request)
		if err != nil {
			util.EPrintf("err in PutInCausal: %v", err)
			return false
		}
		if reply.Vectorclock != nil && reply.Success {
			kvc.vectorclock = reply.Vectorclock
			return reply.Success
		}
		// PutInCausal Failed
		// refresh the target node
		// util.DPrintf("PutInCausal Failed, refresh the target node")
		fmt.Printf("PutInCausal Failed, refresh the target node")
		kvc.kvsId = (kvc.kvsId + 1) % len(kvc.kvservers)
	}
}

var count int32 = 0
var putCount int32 = 0
var getCount int32 = 0
var falseTime int32 = 0

// Test the consistency performance at different read/write ratios
func RequestRatio(cnum int, num int, servers []string, getRatio int, consistencyLevel int) {
	fmt.Printf("servers: %v", servers)
	kvc := KVClient{
		kvservers:   make([]string, len(servers)),
		vectorclock: make(map[string]int32),
	}
	kvc.consistencyLevel = CAUSAL
	for _, server := range servers {
		kvc.vectorclock[server+"1"] = 0
	}
	copy(kvc.kvservers, servers)
	start_time := time.Now()
	for i := 0; i < num; i++ {
		rand.Seed(time.Now().UnixNano())
		key := rand.Intn(100000)
		value := rand.Intn(100000)
		// 写操作
		kvc.PutInCausal("key"+strconv.Itoa(key), "value"+strconv.Itoa(value))
		atomic.AddInt32(&putCount, 1)
		atomic.AddInt32(&count, 1)
		// util.DPrintf("put success")
		for j := 0; j < getRatio; j++ {
			// 读操作
			k := "key" + strconv.Itoa(key)
			v, _ := kvc.GetInCausal(k)
			// if GetInCausal return, it must be success
			atomic.AddInt32(&getCount, 1)
			atomic.AddInt32(&count, 1)
			if v != "" {
				// 查询出了值就输出，屏蔽请求非Leader的情况
				// util.DPrintf("TestCount: ", count, ",Get ", k, ": ", ck.Get(k))
				util.DPrintf("TestCount: %v ,Get %v: %v, VectorClock: %v, getCount: %v, putCount: %v", count, k, v, kvc.vectorclock, getCount, putCount)
				// util.DPrintf("spent: %v", time.Since(start_time))
			}
		}
		// 随机切换下一个节点
		kvc.kvsId = rand.Intn(len(kvc.kvservers)+10) % len(kvc.kvservers)
	}
	fmt.Printf("TestCount: %v, VectorClock: %v, getCount: %v, putCount: %v", count, kvc.vectorclock, getCount, putCount)
	if int(count) == num*cnum*(getRatio+1) {
		fmt.Printf("Task is completed, spent: %v", time.Since(start_time))
		fmt.Printf("falseTimes: %v", falseTime)
	}
}

func main() {
	var ser = flag.String("servers", "", "the Server, Client Connects to")
	var mode = flag.String("mode", "read", "Read or Put and so on")
	var cnums = flag.String("cnums", "1", "Client Threads Number")
	var onums = flag.String("onums", "1", "Client Requests Times")
	var getratio = flag.String("getratio", "1", "Get Times per Put Times")
	var cLevel = flag.Int("consistencyLevel", CAUSAL, "Consistency Level")
	flag.Parse()
	servers := strings.Split(*ser, ",")
	clientNumm, _ := strconv.Atoi(*cnums)
	optionNumm, _ := strconv.Atoi(*onums)
	getRatio, _ := strconv.Atoi(*getratio)
	consistencyLevel := int(*cLevel)

	if clientNumm == 0 {
		fmt.Println("### Don't forget input -cnum's value ! ###")
		return
	}
	if optionNumm == 0 {
		fmt.Println("### Don't forget input -onumm's value ! ###")
		return
	}

	// Request Times = clientNumm * optionNumm
	if *mode == "RequestRatio" {
		for i := 0; i < clientNumm; i++ {
			go RequestRatio(clientNumm, optionNumm, servers, getRatio, consistencyLevel)
		}
	} else {
		fmt.Println("### Wrong Mode ! ###")
		return
	}
	// fmt.Println("count")
	// time.Sleep(time.Second * 3)
	// fmt.Println(count)
	//	return

	// keep main thread alive
	time.Sleep(time.Second * 1200)
}
