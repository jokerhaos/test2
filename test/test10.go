package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func fetchExchangeRate(ctx context.Context, url string, done chan struct{}) {
	defer func() { done <- struct{}{} }()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Failed to create request:", err)
		return
	}

	// 将请求与上下文关联
	req = req.WithContext(ctx)

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Request failed:", err)
		return
	}
	defer resp.Body.Close()

	// 处理响应
	// 这里可以根据实际情况解析和处理响应数据
	fmt.Println("Received response from:", url)
}

func main() {
	// 创建一个父级上下文
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 定义要请求的汇率接口列表
	urls := []string{
		"https://api.exchangerate-api.com/v4/latest/USD",
		"https://api.exchangeratesapi.io/latest?base=USD",
		"https://api.coinbase.com/v2/exchange-rates?currency=USD",
	}
	done := make(chan struct{}, 1)
	// 启动并发请求
	for _, url := range urls {
		go fetchExchangeRate(ctx, url, done)
	}
	select {
	case <-done:
		fmt.Println("请求完成")
	case <-time.After(time.Second * 3):
		fmt.Println("请求超时")
		return
	}

	// 终止所有请求
	cancel()

	// 等待一段时间，以确保所有请求都已终止
	time.Sleep(1 * time.Second)

	fmt.Println("All requests terminated")
}
