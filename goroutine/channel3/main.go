package main

import "fmt"

func main() {
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 1000)
	exitChan := make(chan bool, 4)

	go putNum(intChan)
	for i := 0; i < 4; i++ {
		go putPrimeChan(intChan, primeChan, exitChan)
	}

	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}
		close(primeChan)
	}()

	for {
		res, ok := <-primeChan
		if !ok {
			fmt.Println("全部完成")
			break
		}
		fmt.Println("读取素数", res)
	}
}

func putNum(intChan chan<- int) {
	for i := 0; i < 80; i++ {
		intChan <- i + 1
	}
	close(intChan)
}

func putPrimeChan(intChan <-chan int, primeChan chan<- int, exitChan chan bool) {
	// 取intChan数据
	for {
		num, ok := <-intChan
		if !ok {
			fmt.Println("######################")
			break
		}
		// 判断是不是素数（除1和它本身外，不能被其他的正整数所整除）
		flag := true
		if num == 1 {
			continue
		}
		for i := 2; i < num; i++ {
			if num%i == 0 { // 说明不是素数
				flag = false
				break
			}
		}
		if flag {
			// fmt.Println("放入素数", num)
			primeChan <- num
		}
	}
	exitChan <- true
	fmt.Println("完成工作")
}
