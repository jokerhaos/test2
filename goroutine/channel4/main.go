package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 使用select解决管道读取数据阻塞问题

	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}
	strChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		strChan <- "hello" + strconv.Itoa(i)
	}
	// for range 不关闭管道会阻塞报错 deadlock
	// 实际开发中可能不知道什么时候关闭管道,这时候可以使用select解决
	for {
		select {
		// 注意：这里不会因为管道一直没有关闭，不会一直阻塞二deadlock
		// 会自动下一个case匹配
		case v := <-intChan:
			fmt.Println("intChan", v)
		case v := <-strChan:
			fmt.Println("strChan", v)
		default:
			fmt.Println("没数据可取")
			return
		}
	}

}
