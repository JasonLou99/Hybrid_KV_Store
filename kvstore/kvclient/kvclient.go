package main

/*
	kvclient有三种方式进行通信，http协议、原生tcp协议、grpc通信
	前两种面向wasm runtime调用，后一种是常规的kvs通信模式，用于性能对比
*/
