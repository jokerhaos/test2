package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// context
func init() {
	var v *int
	// v 是一个 int 类型的指针，v 的地址和 v 的值  0xc0000ba018 <nil>
	fmt.Println("v 是一个 int 类型的指针，v 的地址和 v 的值 ", &v, v)
	// 分配给 v 一个指向的变量
	v = new(int)
	// v 是一个 int 类型的指针，v 的地址和 v 的值  0xc0000ba018 0xc000018030 0，此时已经分配给了 v 指针一个指向的变量，但是变量为零值
	fmt.Println("v 是一个 int 类型的指针，v 的地址， v 的值和 v 指向的变量的值 ", &v, v, *v)
	*v = 8
	// v 是一个 int 类型的指针，v 的地址和 v 的值  0xc0000ba018 0xc000018030 8，此时又像这个变量中装填了一个值 8
	fmt.Println("v 是一个 int 类型的指针，v 的地址， v 的值和 v 指向的变量的值 ", &v, v, *v)

	// 整个过程可以理解为给 v 指针指向了一个匿名变量

	// var sm sync.Map
	// sm.Store(

	// ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	// // defer cancel()
	// go func(ctx context.Context) {
	// 	time.Sleep(time.Second * 1)
	// 	fmt.Println("time.After...")
	// 	select {
	// 	case <-time.After(time.Second * 1):
	// 		fmt.Println("time.After...")
	// 	case <-ctx.Done():
	// 		fmt.Println("退出监听协程1", ctx.Err())
	// 		return
	// 		// default:
	// 		// 	fmt.Println("逻辑处理中...")
	// 	}

	// 	time.Sleep(time.Second * 1)
	// 	select {
	// 	case <-ctx.Done():
	// 		fmt.Println("退出监听协程2", ctx.Err())
	// 		return
	// 	default:
	// 		fmt.Println("逻辑处理中...")
	// 	}

	// 	time.Sleep(time.Second * 1)
	// 	select {
	// 	case <-ctx.Done():
	// 		fmt.Println("退出监听协程2", ctx.Err())
	// 		return
	// 	default:
	// 		fmt.Println("逻辑处理中...")
	// 	}

	// }(ctx)
	// time.Sleep(time.Second * 5)

	ch := make(chan int, 1)
	go func() {
		time.Sleep(time.Second * 1)
	}()
	ch <- 10
	ch <- 10
	fmt.Println("发送成功")

	// 设置超时时间100ms
	ctx, _ := context.WithTimeout(context.Background(), 100*time.Millisecond)

	// 构建一个HTTP请求
	req, _ := http.NewRequest(http.MethodGet, "https://www.baidu.com/", nil)
	// 把ctx信息传进去
	req = req.WithContext(ctx)

	client := &http.Client{}
	// 向百度发送请求
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Request failed:", err)
		return
	}
	fmt.Println("Response received, status code:", res.StatusCode)

}

// select creator_uid from image where img_suffix="png" order by created_at limit 10
