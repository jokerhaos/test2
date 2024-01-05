package main

import (
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
	// "github.com/gomodule/redigo/redis"
)

var pool *redis.Pool

func init() {
	pool = &redis.Pool{
		MaxIdle:     8,                 // 最大空闲数
		MaxActive:   0,                 // 最大连接数 0无限制
		IdleTimeout: 100 * time.Second, // 最大空闲时间
		Wait:        true,              // 超过连接数后是否等待
		Dial: func() (redis.Conn, error) { // 初始化连接
			return redis.Dial("tcp", "127.0.0.1:6379", redis.DialPassword("123456"))
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}

func main() {
	// 从连接池取出一个连接
	conn := pool.Get()
	defer conn.Close() // 关闭连接池，一旦关闭就不能取数据了

	res, err := conn.Do("SET", "sssss5", 1, "NX", "EX", 300)
	fmt.Println(res, err)
}
