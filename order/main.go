package main

import "fmt"

func main() {
	fbn := fbn(10)
	fmt.Println(fbn)
	bubble := [6]uint64{123, 4123, 2543, 1, 1234, 5}
	bubbleSort(&bubble)
}

// 斐波那契
func fbn(n int) []uint64 {
	fbnSlice := make([]uint64, n)
	fbnSlice[0] = 1
	fbnSlice[1] = 1
	for i := 2; i < n; i++ {
		fbnSlice[i] = fbnSlice[i-1] + fbnSlice[i-2]
	}
	return fbnSlice
}

// 冒泡排序
func bubbleSort(arr *[6]uint64) {
	fmt.Println("冒泡排序前", *arr)
	len := len(*arr)
	for i := 0; i < len-1; i++ {
		for j := 0; j < len-i-1; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				// 交换
				(*arr)[j] = (*arr)[j] ^ (*arr)[j+1]
				(*arr)[j+1] = (*arr)[j] ^ (*arr)[j+1]
				(*arr)[j] = (*arr)[j] ^ (*arr)[j+1]
			}
		}

	}
	fmt.Println("冒泡排序后", *arr)
}
