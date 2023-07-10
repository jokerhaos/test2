package main

import (
	"fmt"
	"sort"
	"time"

	"github.com/samber/lo"
)

// 添加一下自定义 map 类型
type MyMap map[string]interface{}

// 格式化时间
func (my MyMap) formTime(names ...string) MyMap {
	for k, value := range my {
		if lo.Contains(names, k) {
			// 时间格式化
			switch v := value.(type) {
			case string:
				t, err := time.Parse(time.RFC3339, v)
				if err == nil {
					my[k+"_form"] = t.Format("2006-01-02 15:04:05")
				}
			case time.Time:
				my[k+"_form"] = v.Format("2006-01-02 15:04:05")
			}
		}
	}

	return my
}

// 过滤某部分 key
func (my MyMap) filterByKey(keys ...string) MyMap {
	for k, _ := range my {
		if lo.Contains(keys, k) {
			delete(my, k)
		}
	}
	return my
}

// 获取部分key
func (my MyMap) getByKey(keys ...string) MyMap {
	m := make(MyMap, 0)
	for k, v := range my {
		if lo.Contains(keys, k) {
			m[k] = v
		}
	}
	return my
}

// 按照 key 的顺序循环
func (my MyMap) iterateInOrder(fn func(k string, v interface{})) {
	keys := make([]string, 0, len(my))
	for k := range my {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fn(k, my[k])
	}
}

// 按照 key 的倒序循环
func (my MyMap) iterateInReverseOrder(fn func(k string, v interface{})) {
	keys := make([]string, 0, len(my))
	for k := range my {
		keys = append(keys, k)
	}
	sort.Sort(sort.Reverse(sort.StringSlice(keys)))
	for _, k := range keys {
		fn(k, my[k])
	}
}

// 等于某个 key 时回调
func (my MyMap) onKeyMatch(key string, callback func(value interface{})) {
	if v, ok := my[key]; ok {
		callback(v)
	}
}

func main() {
	my := MyMap{
		"time1": "2023-07-03T10:15:30Z",
		"time2": time.Now(),
		"key1":  "value1",
		"key2":  "value2",
	}

	// 格式化时间
	my.formTime("time1", "time2")
	fmt.Println("------------格式化时间-----------")
	fmt.Println(my)

	// 过滤某个 key
	filtered := my.filterByKey("key1")
	fmt.Println("------------过滤key1-----------")
	fmt.Println(filtered)

	fmt.Println("------------循序map遍历-----------")
	// 按照 key 的顺序循环
	my.iterateInOrder(func(k string, v interface{}) {
		fmt.Println(k, v)
	})

	fmt.Println("------------倒序map遍历-----------")
	// 按照 key 的倒序循环
	my.iterateInReverseOrder(func(k string, v interface{}) {
		fmt.Println(k, v)
	})

	fmt.Println("-----------------------")
	// 等于某个 key 时回调
	my.onKeyMatch("time2", func(value interface{}) {
		fmt.Println("Value:", value)
	})
}
