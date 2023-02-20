package config

type Log struct {
	Option string
	Key    string
	Value  string
}

// 对外提供kv服务的地址
// 首字母大写，表示public
var Address string = "192.168.10.120:7088"
