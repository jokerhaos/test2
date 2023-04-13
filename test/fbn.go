package main

import "fmt"

func main() {
	n := 5 // 台阶数
	ways := climbStairs(n)
	fmt.Printf("对于 %d 阶台阶，共有 %d 种上法\n", n, ways)
}

func climbStairs(n int) int {
	if n <= 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	// if n == 3 {
	// 	return 4
	// }
	return climbStairs(n-1) + climbStairs(n-2)
}

func climbStairs2(n int) int {
	if n <= 1 {
		return 1
	}

	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	fmt.Println(dp)
	return dp[n]
}
