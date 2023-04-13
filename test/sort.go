package main

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

type Test struct {
	B string  `json:"b"`
	A int     `json:"a"`
	D string  `json:"d"`
	C float64 `json:"c"`
}

func main() {
	t := &Test{
		A: 1,
		D: "d",
		B: "b",
		C: 2,
	}

	// 使用反射获取结构体字段
	val := reflect.ValueOf(t).Elem()
	typ := val.Type()

	// 构建存储字段信息的切片
	fields := make([]string, val.NumField())

	// 遍历结构体字段，将字段名和对应的值存储到切片中
	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		if field.PkgPath != "" {
			// 跳过未导出字段
			continue
		}
		fieldName := field.Tag.Get("json")                        // 通过标签获取字段名
		fieldValue := fmt.Sprintf("%v", val.Field(i).Interface()) // 将字段值转换为字符串
		fields[i] = fmt.Sprintf("%s=%s", fieldName, fieldValue)
	}

	// 按照字段名排序
	sort.Strings(fields)

	// 将排序后的字段拼接成字符串
	result := strings.Join(fields, "&")

	fmt.Println(result)
}
