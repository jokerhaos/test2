// 可以引⼊的库和版本相关请参考 “环境说明”
// 必须定义⼀个 包名为 `main` 的包

package main

import (
	"fmt"
	"math"
)

type Node struct {
	Value    int
	Children []*Node
}

// 给出一个数字在map数组中找到最相近的数字
func solution(ranks map[int]int) int {
	ranks[1] = 93
	ranks[10] = 55
	ranks[15] = 30
	ranks[20] = 19
	ranks[23] = 11
	ranks[30] = 2
	ho := 19

	gap := 0
	prevGap := 0
	result := 0

	for rank, honor := range ranks {
		gap = int(math.Abs(float64(honor - ho)))
		if gap <= prevGap {
			result = rank
		}
		prevGap = gap
	}
	return result
}

func main() {
	ranks := map[int]int{}
	rank := solution(ranks)
	fmt.Println(rank)
	// var s string
	// s = "abcdefg"
	// for i, char := range s {
	// 	fmt.Println(string(char), i)
	// }
}
