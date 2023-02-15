package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

var pool *redis.Pool

func init() {
	pool = &redis.Pool{
		MaxIdle:     8,   // 最大空闲数
		MaxActive:   0,   // 最大连接数 0无限制
		IdleTimeout: 100, // 最大空闲时间
		Dial: func() (redis.Conn, error) { // 初始化连接
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
	}
}

func main() {
	// 从连接池取出一个连接
	connPool := pool.Get()
	// pool.Close() // 关闭连接池，一旦关闭就不能取数据了
	result, err := redis.String(connPool.Do("name"))
	fmt.Println("name: ", result)

	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis.Dial err=", err)
		return
	}
	fmt.Println("redis连接成功")

	defer conn.Close()

	// 存
	_, err = conn.Do("Set", "name", "joker")
	if err != nil {
		fmt.Println("redis Set err=", err)
		return
	}

	// 取
	r, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("redis Get err=", err)
		return
	}

	fmt.Println("name: ", r)

}
