package main

import (
	"fmt"
)

func QuickSort(arr []int, low, high int) {
	if low < high {
		// 选择枢轴元素，这里简单选择数组的最后一个元素作为枢轴
		pivot := arr[high]
		i := low - 1
		for j := low; j < high; j++ {
			if arr[j] <= pivot {
				i++
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
		arr[i+1], arr[high] = arr[high], arr[i+1]

		// 递归对枢轴左右两侧的子数组进行排序
		QuickSort(arr, low, i)
		QuickSort(arr, i+2, high)
	}
}

func QuickSort2(array []int, left int, right int) {
	l := left
	r := right
	// pivot 表示中轴
	pivot := array[(left+right)/2]
	//for循环的目标是将比pivot小的数放到左边，比pivot大的数放到右边
	for l < r {
		//从pivot左边找到大于等于pivot的值
		for array[l] < pivot {
			l++
		}
		//从pivot右边找到大于等于pivot的值
		for array[r] > pivot {
			r--
		}
		//交换位置
		array[l], array[r] = array[r], array[l]
		fmt.Println(array)
		//优化
		if l == r {
			l++
			r--
		}
		//向左递归
		if left < r {
			QuickSort2(array, left, r)
		}
		//向左递归
		if right > l {
			QuickSort2(array, l, right)
		}
	}
}

func main() {
	arr := []int{5, 1, 9, 3, 7, 6, 8, 2, 4}
	fmt.Println("Original array:", arr)
	QuickSort2(arr, 0, len(arr)-1)
	// quickSort(arr, 0, len(arr)-1)
	fmt.Println("Sorted array:", arr)
}
