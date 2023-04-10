package main

import (
	"fmt"
	"sync"
	"time"
)

// 问一道思考题：如何并发100个任务，但是同一时间最多运行的10个任务
func init() {
	count := 10

	num := 1000
	wg := sync.WaitGroup{}

	c := make(chan struct{}, count)
	for i := 0; i < num; i++ {
		wg.Add(1)
		c <- struct{}{}
		go func(j int) {
			defer wg.Done()
			fmt.Println(j)
			<-c
			time.Sleep(time.Second * 10)
		}(i)
	}
	wg.Wait()

}
