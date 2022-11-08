package main

import (
	"fmt"
	"strconv"
	"time"
)

// 主线程执行完毕，协程没执行完也会退出！
func main() {
	// 协程中每隔一秒执行
	go test()
	// 主线程中每隔一秒执行
	for i := 0; i < 10; i++ {
		fmt.Println("main hello,world" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func test() {
	for i := 0; i < 10; i++ {
		fmt.Println("go hello,world" + strconv.Itoa(i))
		time.Sleep(time.Second * 2)
	}
}
