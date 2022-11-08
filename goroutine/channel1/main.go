package main

import "fmt"

// 管道先进先出
func main() {
	test()
}

func test() {
	// var intChan chan int

	// 声明只写
	// var chan2 chan<- int

	// 声明只读
	// var chan3 <-chan int

	intChan := make(chan int, 53)

	fmt.Printf("%T,%v\n", intChan, intChan)

	// 塞入数据
	intChan <- 10
	num := 200
	intChan <- num
	intChan <- 30
	intChan <- 40
	intChan <- 50
	intChan <- 60

	fmt.Printf("intChan len = %v , intChan cap = %v\n", len(intChan), cap(intChan))

	// 取出数据
	num2 := <-intChan
	fmt.Printf("num2 = %v\n", num2)
	fmt.Printf("intChan len = %v , intChan cap = %v\n", len(intChan), cap(intChan))

	// 关闭管道之后就不可以写入数据，但是可以取数据！
	close(intChan)
	// intChan <- 40
	// fmt.Printf("intChan len = %v , intChan cap = %v\n", len(intChan), cap(intChan))
	num3 := <-intChan
	fmt.Printf("num3 = %v\n", num3)
	fmt.Printf("intChan len = %v , intChan cap = %v\n", len(intChan), cap(intChan))

	// 遍历管道
	intChan2 := make(chan int, 100)
	for i := 0; i < 100; i++ {
		intChan2 <- i * 2
	}
	fmt.Printf("intChan2 len = %v , intChan cap = %v\n", len(intChan2), cap(intChan2))

	close(intChan2)
	// 遍历必须先关闭管道！！！
	for v := range intChan2 {
		fmt.Printf("v = %v\n", v)
	}
	fmt.Printf("intChan2 len = %v , intChan cap = %v\n", len(intChan2), cap(intChan2))

}
