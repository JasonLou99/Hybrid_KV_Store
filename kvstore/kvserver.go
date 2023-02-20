package main

import (
	"hybrid_kv_store/config"
	"hybrid_kv_store/persister"
	"hybrid_kv_store/util"
	"sync"
)

type KVServer struct {
	peer        []KVServer
	address     string
	log         config.Log
	vectorclock sync.Map
	persister   *persister.Persister
}

func (kvs *KVServer) Get(k string) string {
	// TODO: implement
	return ""
}

func (kvs *KVServer) Put(k string, v string) {
	// TODO: implement

}

func MakeKVServer() *KVServer {
	util.IPrintf("Make KVServer %s... ", config.Address)
	kvs := new(KVServer)
	kvs.persister = new(persister.Persister)
	kvs.persister.Init("data")
	kvs.address = config.Address
	return kvs
}
