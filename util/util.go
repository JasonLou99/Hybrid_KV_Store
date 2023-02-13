package util

import (
	"log"
	"sync"
)

// Debugging
const Debug = true

func DPrintf(format string, a ...interface{}) (n int, err error) {
	if Debug {
		log.SetFlags(log.Ldate | log.Lmicroseconds)
		log.Printf(format, a...)
	}
	return
}

/*
sync.Map 相关的函数
*/
func Len(vc sync.Map) int {
	count := 0
	vc.Range(func(key, value interface{}) bool {
		count++
		return true
	})
	return count
}
func BecomeMap(vc sync.Map) map[string]int32 {
	res := map[string]int32{}
	vc.Range(func(key, value interface{}) bool {
		res[key.(string)] = value.(int32)
		return true
	})
	return res
}
func BecomeSyncMap(argMap map[string]int32) sync.Map {
	var res sync.Map
	for key, value := range argMap {
		res.Store(key, value)
	}
	return res
}

// 判断vectorClock是否更大（key都有，并且value>=other.value）
func IsUpper(vectorClock sync.Map, vc sync.Map) bool {
	DPrintf("IsUpper(): vectorClock: %v, arg_vc: %v", BecomeMap(vectorClock), BecomeMap(vc))
	if Len(vc) == 0 {
		return true
	}
	if Len(vectorClock) < Len(vc) {
		return false
	} else {
		res := true
		vc.Range(func(k, v interface{}) bool {
			value, ok := vectorClock.Load(k)
			if ok {
				if v.(int32) <= value.(int32) {
					return true
				} else {
					res = false
					// 遍历终止
					return false
				}
			}
			// vectorClock中没有该key
			res = false
			// 遍历终止
			return false
		})
		return res
	}
}
