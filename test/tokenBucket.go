package main

import (
	"fmt"
	"sync"
	"time"
)

type TokenBucket struct {
	rate       int
	capacity   int
	tokens     int
	lastRefill time.Time
	mu         sync.Mutex // 互斥锁
}

func NewTokenBucket(rate, capacity int) *TokenBucket {
	tb := &TokenBucket{
		rate:       rate,
		capacity:   capacity,
		tokens:     capacity,
		lastRefill: time.Now(),
	}
	return tb
}

func (tb *TokenBucket) Take() bool {
	tb.mu.Lock()
	defer tb.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(tb.lastRefill)

	refillCount := int(elapsed.Seconds()) * tb.rate
	if refillCount > 0 {
		tb.tokens = tb.tokens + refillCount
		if tb.tokens > tb.capacity {
			tb.tokens = tb.capacity
		}
		tb.lastRefill = now
	}

	if tb.tokens > 0 {
		tb.tokens--
		return true
	}

	return false
}

func main() {
	tb := NewTokenBucket(5, 10)

	// 使用 WaitGroup 来等待所有 goroutine 完成
	var wg sync.WaitGroup
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			if tb.Take() {
				fmt.Println("Request", i, "is processed.")
			} else {
				fmt.Println("Request", i, "is rejected.")
			}
		}(i)
	}

	// 等待所有 goroutine 完成
	wg.Wait()
}
