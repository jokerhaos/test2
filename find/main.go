package main

import "fmt"

func main() {
	arr := [6]uint64{1, 3, 4, 5, 6, 12}
	BinaryFind(&arr, 0, uint64(len(arr)-1), 6)
}

func BinaryFind(arr *[6]uint64, left uint64, right uint64, value uint64) {
	if left > right {
		fmt.Println("找不到")
		return
	}
	middle := (left + right) / 2
	if (*arr)[middle] > value {
		BinaryFind(arr, left, middle-1, value)
	} else if (*arr)[middle] < value {
		BinaryFind(arr, middle+1, right, value)
	} else {
		fmt.Println("找到了", middle)
	}
}
