package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"hybrid_kv_store/config"
	"hybrid_kv_store/lattices"
	"hybrid_kv_store/persister"
	"hybrid_kv_store/rpc/causalrpc"
	"hybrid_kv_store/rpc/kvrpc"
	"hybrid_kv_store/util"
	"math/rand"
	"net"
	"strings"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

type KVServer struct {
	peers           []string
	address         string
	internalAddress string // internal address for communication between nodes
	latency         int    // Simulation of geographical delay
	logs            []config.Log
	vectorclock     sync.Map
	persister       *persister.Persister
	db              sync.Map // memory database
	// causalEntity *causal.CausalEntity
}

type ValueTimestamp struct {
	value     string
	timestamp int64
}

// this method is used to execute the command from client with causal consistency
func (kvs *KVServer) startInCausal(command interface{}, vcFromClientArg map[string]int32, timestampFromClient int64) bool {
	vcFromClient := util.BecomeSyncMap(vcFromClientArg)
	newLog := command.(config.Log)
	util.DPrintf("Log in Start(): %v ", newLog)
	util.DPrintf("vcFromClient in Start(): %v", vcFromClient)
	if newLog.Option == "Put" {
		isUpper := util.IsUpper(vcFromClient, kvs.vectorclock)
		if isUpper {
			vt, ok := kvs.db.Load(newLog.Key)
			vt2 := vt.(ValueTimestamp)
			if ok && vt2.timestamp > timestampFromClient {
				// the value in the db is newer than the value in the client
				return false
			}
			// update vector clock
			kvs.vectorclock = vcFromClient
			val, _ := kvs.vectorclock.Load(kvs.internalAddress)
			kvs.vectorclock.Store(kvs.internalAddress, val.(int32)+1)
			// init MapLattice for sending to other nodes
			ml := lattices.MapLattice{
				Key: newLog.Key,
				Vl: lattices.ValueLattice{
					Log:         newLog,
					VectorClock: util.BecomeMap(kvs.vectorclock),
				},
			}
			data, _ := json.Marshal(ml)
			args := &causalrpc.AppendEntriesInCausalRequest{
				MapLattice: data,
			}
			// async sending to other nodes
			for i := 0; i < len(kvs.peers); i++ {
				if kvs.peers[i] != kvs.address {
					go kvs.sendAppendEntriesInCausal(kvs.peers[i], args)
				}
			}
			// update value in the db and persist
			kvs.logs = append(kvs.logs, newLog)
			kvs.db.Store(newLog.Key, &ValueTimestamp{value: newLog.Value, timestamp: time.Now().UnixMilli()})
			kvs.persister.Put(newLog.Key, newLog.Value)
			return true
		} else {
			return false
		}
	} else if newLog.Option == "Get" {
		return util.IsUpper(kvs.vectorclock, vcFromClient)
	}
	util.DPrintf("here is Start() in Causal: log command option is false")
	return false
}

func (kvs *KVServer) GetInCasual(ctx context.Context, in *kvrpc.GetInCasualRequest) (*kvrpc.GetInCasualResponse, error) {
	util.DPrintf("GetInCasual %s", in.Key)
	getInCausalResponse := new(kvrpc.GetInCasualResponse)
	op := config.Log{
		Option: "Get",
		Key:    in.Key,
		Value:  "",
	}
	ok := kvs.startInCausal(op, in.Vectorclock, in.Timestamp)
	if ok {
		vt, _ := kvs.db.Load(in.Key)
		valueTimestamp := vt.(*ValueTimestamp)
		// compare timestamp
		if valueTimestamp.timestamp > in.Timestamp {
			getInCausalResponse.Value = ""
			getInCausalResponse.Success = false
		}
		// getInCausalResponse.Value = string(kvs.persister.Get(in.Key))
		getInCausalResponse.Value = valueTimestamp.value
		getInCausalResponse.Success = true
	} else {
		getInCausalResponse.Value = ""
		getInCausalResponse.Success = false
	}
	return getInCausalResponse, nil
}

func (kvs *KVServer) PutInCasual(ctx context.Context, in *kvrpc.PutInCasualRequest) (*kvrpc.PutInCasualResponse, error) {
	util.DPrintf("PutInCasual %s %s", in.Key, in.Value)
	putInCausalResponse := new(kvrpc.PutInCasualResponse)
	op := config.Log{
		Option: "Put",
		Key:    in.Key,
		Value:  in.Value,
	}
	ok := kvs.startInCausal(op, in.Vectorclock, in.Timestamp)
	if !ok {
		util.DPrintf("PutInCasual: StartInCausal Failed key=%s value=%s, Because vcFromClient < kvs.vectorclock", in.Key, in.Value)
		putInCausalResponse.Success = false
	} else {
		putInCausalResponse.Success = true
	}
	putInCausalResponse.Vectorclock = util.BecomeMap(kvs.vectorclock)
	return putInCausalResponse, nil
}

func (kvs *KVServer) AppendEntriesInCausal(ctx context.Context, in *causalrpc.AppendEntriesInCausalRequest) (*causalrpc.AppendEntriesInCausalResponse, error) {
	util.DPrintf("AppendEntriesInCausal %v", in)
	var mlFromOther lattices.MapLattice
	json.Unmarshal(in.MapLattice, &mlFromOther)
	vcFromOther := util.BecomeSyncMap(mlFromOther.Vl.VectorClock)
	ok := util.IsUpper(kvs.vectorclock, vcFromOther)
	if !ok {
		// Append the log to the local log
		kvs.logs = append(kvs.logs, mlFromOther.Vl.Log)
		kvs.db.Store(mlFromOther.Key, &ValueTimestamp{value: mlFromOther.Vl.Log.Value, timestamp: time.Now().UnixMilli()})
		kvs.persister.Put(mlFromOther.Key, mlFromOther.Vl.Log.Value)
	}
	return &causalrpc.AppendEntriesInCausalResponse{}, nil
}

func (kvs *KVServer) RegisterKVServer(address string) {
	util.DPrintf("RegisterKVServer: %s", address)
	for {
		lis, err := net.Listen("tcp", address)
		if err != nil {
			util.FPrintf("failed to listen: %v", err)
		}
		grpcServer := grpc.NewServer()
		kvrpc.RegisterKVServer(grpcServer, kvs)
		reflection.Register(grpcServer)
		if err := grpcServer.Serve(lis); err != nil {
			util.FPrintf("failed to serve: %v", err)
		}
	}
}

func (kvs *KVServer) RegisterCausalServer(address string) {
	util.DPrintf("RegisterCausalServer: %s", address)
	for {
		lis, err := net.Listen("tcp", address)
		if err != nil {
			util.FPrintf("failed to listen: %v", err)
		}
		grpcServer := grpc.NewServer()
		causalrpc.RegisterCAUSALServer(grpcServer, kvs)
		reflection.Register(grpcServer)
		if err := grpcServer.Serve(lis); err != nil {
			util.FPrintf("failed to serve: %v", err)
		}
	}
}

// s0 --> other servers
func (kvs *KVServer) sendAppendEntriesInCausal(address string, args *causalrpc.AppendEntriesInCausalRequest) (*causalrpc.AppendEntriesInCausalResponse, bool) {
	util.DPrintf("here is sendAppendEntriesInCausal() ---------> ", address)
	// 随机等待，模拟延迟
	time.Sleep(time.Millisecond * time.Duration(kvs.latency+rand.Intn(25)))
	// conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		util.EPrintf("sendAppendEntriesInCausal did not connect: %v", err)
	}
	defer conn.Close()
	client := causalrpc.NewCAUSALClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	reply, err := client.AppendEntriesInCausal(ctx, args)
	if err != nil {
		util.EPrintf("sendAppendEntriesInCausal could not greet: ", err, address)
		return reply, false
	}
	return reply, true
}

func MakeKVServer(peers []string) *KVServer {
	util.IPrintf("Make KVServer %s... ", config.Address)
	kvs := new(KVServer)
	kvs.persister = new(persister.Persister)
	kvs.persister.Init("db")
	kvs.address = config.Address
	kvs.internalAddress = config.InternalAddress
	kvs.peers = peers
	// init vectorclock: { "192.168.10.120:30881":0, "192.168.10.121:30881":0, ... }
	for i := 0; i < len(peers); i++ {
		kvs.vectorclock.Store(peers[i], 0)
	}
	return kvs
}

func main() {
	// peers inputed by command line
	var peers_arg = flag.String("peers", "", "Input Your Peers")
	flag.Parse()
	peers := strings.Split(*peers_arg, ",")
	fmt.Println(peers)
	kvs := MakeKVServer(peers)
	go kvs.RegisterKVServer(kvs.address)
	go kvs.RegisterCausalServer(kvs.internalAddress)
	// server run for 20min
	time.Sleep(time.Second * 1200)
}
