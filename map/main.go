package main

import (
	"fmt"
	"sort"
)

type Stu struct {
	Name    string
	Age     int
	Address string
}

// map 遍历是无序的
func main() {
	createMap()

	map1 := make(map[int]int, 10)
	map1[10] = 100
	map1[1] = 200
	map1[2] = 300

	var keys []int
	for k, v := range map1 {
		fmt.Println(v)
		keys = append(keys, k)
	}

	// 通过key排序循环输出
	sort.Ints(keys)

	fmt.Println(keys)
	fmt.Println("-------------排序-------------")
	for i := 0; i < len(keys); i++ {
		fmt.Println(map1[keys[i]])
	}

	Curd()

	maps := make(map[int]*Stu)
	maps[0] = &Stu{
		Name:    "joker",
		Age:     18,
		Address: "湖南",
	}
	maps[1] = &Stu{
		Name:    "tom",
		Age:     10,
		Address: "美国佬",
	}
	// map是引用传递
	fmt.Println(*maps[0])
	modify(maps)
	fmt.Println(*maps[0])

}

func createMap() {
	var map1 map[string]string
	map1 = make(map[string]string, 10)
	map1["1"] = "第一种定义"
	fmt.Println(map1)

	map2 := make(map[string]string)
	map2["2"] = "第二种定义"
	fmt.Println(map2)

	map3 := map[string]string{
		"3": "第三种定义",
	}
	fmt.Println(map3)

}

func Curd() {
	// r
	map1 := make(map[string]string)
	map1["a"] = "ab"
	map1["b"] = "ab2"
	fmt.Println(map1)

	// c
	v, ok := map1["a"]
	if ok {
		fmt.Println("找到了", v)
	} else {
		fmt.Println("没有找到了")
	}

	// u
	map1["a"] = "update"
	fmt.Println(map1)

	// d
	delete(map1, "b")
	fmt.Println(map1)

}

func modify(map1 map[int]*Stu) {
	stu := map1[0]
	(*stu).Age = 20
}
