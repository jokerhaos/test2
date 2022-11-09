package main

import (
	"fmt"
	"strconv"
	"sync"
)

// 1.1.1. sync.WaitGroup
// 在代码中生硬的使用time.Sleep肯定是不合适的，Go语言中可以使用sync.WaitGroup来实现并发任务的同步。 sync.WaitGroup有以下几个方法：

// 方法名	功能
// (wg * WaitGroup) Add(delta int)	计数器+delta
// (wg *WaitGroup) Done()	计数器-1
// (wg *WaitGroup) Wait()	阻塞直到计数器变为0
var wg sync.WaitGroup

// 1.1.2. sync.Once
// 说在前面的话：这是一个进阶知识点。
// 在编程的很多场景下我们需要确保某些操作在高并发的场景下只执行一次，例如只加载一次配置文件、只关闭一次通道等。
// Go语言中的sync包中提供了一个针对只执行一次场景的解决方案–sync.Once。
// sync.Once只有一个Do方法，其签名如下：

func main() {
	defer wg.Wait()

	intChan := make(chan int, 50)

	go writeData(intChan)
	go readData(intChan)

	fmt.Println("任务完成")

	go Icon("aaaa")
	go Icon("vvvv")
	go testMap()
	fmt.Println(m)
	fmt.Println(m2)

}

// write
func writeData(intChan chan int) {
	wg.Add(1)
	defer wg.Done()
	for i := 0; i < 5; i++ {
		// 放入数据
		intChan <- i + 1
		// fmt.Println("写入数据", i+1)
		// time.Sleep(time.Millisecond)
	}
	close(intChan)
}

// read
func readData(intChan chan int) {
	wg.Add(1)
	defer wg.Done()
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Println("读取到数据", v)
	}
}

var icons map[string]string

var loadIconsOnce sync.Once

func loadIcons() {
	icons = map[string]string{
		"left":  "left.png",
		"up":    "up.png",
		"right": "right.png",
		"down":  "down.png",
	}
	fmt.Println(icons)
}

func Icon(name string) string {
	wg.Add(1)
	defer wg.Done()
	// // Icon 被多个goroutine调用时不是并发安全的
	// if icons == nil {
	// 	loadIcons()
	// }
	// return icons[name]

	// Icon 是并发安全的
	loadIconsOnce.Do(loadIcons)
	return icons[name]
}

// Go语言中内置的map不是并发安全的。
var m = make(map[string]int)

func get(key string) int {
	return m[key]
}

func set(key string, value int) {
	m[key] = value
}

// Go语言的sync包中提供了一个开箱即用的并发安全版map–sync.Map。
// 开箱即用表示不用像内置的map一样使用make函数初始化就能直接使用。
// 同时sync.Map内置了诸如Store、Load、LoadOrStore、Delete、Range等操作方法。
var m2 = sync.Map{}

func testMap() {
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			// 并发安全
			m2.Store(key, n)
			value, _ := m2.Load(key)
			// 不安全
			// set(key, n)

			fmt.Printf("k=:%v,v:=%v\n", key, value)
			wg.Done()
		}(i)
	}
}
