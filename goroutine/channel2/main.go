package main

import (
	"fmt"
)

func main() {
	intChan := make(chan int, 50)
	exitChan := make(chan bool, 1)
	go writeData(intChan)
	go readData(intChan, exitChan)
	for {
		v, ok := <-exitChan
		if ok && v {
			fmt.Println("任务完成")
			break
		}
	}
}

// write
func writeData(intChan chan int) {
	for i := 0; i < 50; i++ {
		// 放入数据
		intChan <- i + 1
		fmt.Println("写入数据", i+1)
		// time.Sleep(time.Millisecond)
	}
	close(intChan)
}

// read
func readData(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Println("读取到数据", v)
	}
	exitChan <- true
	close(exitChan)
}
