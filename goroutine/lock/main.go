package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	for i := 1; i <= 20; i++ {
		go test(i)
	}
	time.Sleep(time.Second * 3)
	lock.Lock()
	for i, v := range myMap {
		fmt.Printf("key=%v,value=%v\r\n", i, v)
	}
	lock.Unlock()
}

var myMap = make(map[int]uint64, 10)
var lock sync.Mutex

func test(n int) {
	res := uint64(1)
	for i := 1; i <= n; i++ {
		res *= uint64(i)
	}
	// 加锁，底层会进行资源争夺 myMap
	lock.Lock()
	myMap[n] = res
	lock.Unlock()
}
