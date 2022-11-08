package main

import (
	"fmt"
	"runtime"
)

func main() {
	num := runtime.NumCPU()
	fmt.Println("系统一共有", num, "个cpu")

	// 可以设置使用多个cpu
	runtime.GOMAXPROCS(num / 2)

}
