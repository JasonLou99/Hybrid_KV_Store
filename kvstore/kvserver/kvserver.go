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

	// causalEntity *causal.CausalEntity
}

// // this method is used to execute the command from client with causal consistency
func (kvs *KVServer) startInCausal(command interface{}, vcFromClientArg map[string]int32) bool {
	vcFromClient := util.BecomeSyncMap(vcFromClientArg)
	newLog := command.(config.Log)
	util.DPrintf("Log in Start(): %v ", newLog)
	util.DPrintf("vcFromClient in Start(): %v", vcFromClient)
	if newLog.Option == "Put" {
		isUpper := util.IsUpper(vcFromClient, kvs.vectorclock)
		if isUpper {
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
			kvs.logs = append(kvs.logs, newLog)
			kvs.persister.Put(newLog.Key, newLog.Value)
			// kvs.applyCh <- 1
			// util.DPrintf("applyCh unread buffer: %v", len(ce.applyCh))
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
	// getInCausalResponse := new(kvrpc.GetInCasualResponse)
	// op := config.Log{
	// 	Option: "get",
	// 	Key:    in.Key,
	// 	Value:  "",
	// }

	return &kvrpc.GetInCasualResponse{}, nil
}

func (kvs *KVServer) PutInCasual(ctx context.Context, in *kvrpc.PutInCasualRequest) (*kvrpc.PutInCasualResponse, error) {
	return &kvrpc.PutInCasualResponse{}, nil
}

func (kvs *KVServer) AppendEntriesInCausal(ctx context.Context, in *causalrpc.AppendEntriesInCausalRequest) (*causalrpc.AppendEntriesInCausalResponse, error) {
	util.DPrintf("AppendEntriesInCausal")
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
	fmt.Println("here is sendAppendEntriesInCausal() ---------> ", address)
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
	kvs.internalAddress = config.Address + "1"
	kvs.peers = peers
	// init vectorclock: { "192.168.10.120:7081":0, "192.168.10.121:7081":0, ... }
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
