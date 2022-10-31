package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	test1()
}

// 生成10个1-100的随机整数,倒叙打印，求平均值、最大值、最小值
func test1() {
	rand.Seed(time.Now().UnixNano())
	// 生成随机数
	arr := [10]uint64{}
	for i := 0; i < 10; i++ {
		arr[i] = uint64(rand.Intn(99) + 1)
	}
	fmt.Println("倒叙前", arr)
	len := len(arr)
	// 倒叙打印
	for i := 0; i < len/2; i++ {
		// 当前位置和最后一个换位置
		arr[i] = arr[i] ^ arr[len-i-1]
		arr[len-i-1] = arr[i] ^ arr[len-i-1]
		arr[i] = arr[i] ^ arr[len-i-1]
	}
	fmt.Println("倒叙后", arr)
}
