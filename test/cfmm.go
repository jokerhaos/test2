package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 根据时间生成随机种子
	rand.Seed(time.Now().UnixNano())

	// 执行2142万次随机生成算法
	count := 21420000
	var result []int
	for i := 1; i <= count; i++ {
		if count >= 100000 && count%100000 == 0 {
			fmt.Printf("正在执行，当前执行了 %d 次了\n", i)
		}
		result = generateLotteryNumbers()
		// 处理生成的号码，可以根据需求进行相应的操作
		_ = result
	}

	fmt.Printf("执行完毕，生成了 %d 组号码\n", count)
	fmt.Printf("执行完毕，生成了 %d 组号码\n", result)
}

// 生成大乐透号码
func generateLotteryNumbers() []int {
	// 生成红色球号码
	redBalls := generateBalls(1, 35, 5)

	// 生成蓝色球号码
	blueBalls := generateBalls(1, 12, 2)

	// 合并红色球和蓝色球号码
	numbers := append(redBalls, blueBalls...)

	return numbers
}

// 生成指定范围内的随机号码
func generateBalls(min, max, count int) []int {
	balls := make([]int, 0)
	for len(balls) < count {
		// 生成随机号码
		ball := rand.Intn(max-min+1) + min

		// 检查是否已存在该号码
		exists := false
		for _, b := range balls {
			if b == ball {
				exists = true
				break
			}
		}

		// 如果号码不存在，则添加到列表中
		if !exists {
			balls = append(balls, ball)
		}
	}

	return balls
}
