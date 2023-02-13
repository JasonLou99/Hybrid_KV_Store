package main

import "hybrid_kv_store/config"

type KVServer struct {
	peer    []KVServer
	address string
	log     config.Log
}

func (kvs *KVServer) Get(k string) string {
	return _
}

func (kvs *KVServer) put(k string, v string) {

}
