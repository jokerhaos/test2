package main

import (
	"fmt"
	"reflect"
)

func IsNull(i interface{}) {
	v := reflect.ValueOf(i)
	if v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
		if v.IsNil() {
			fmt.Println("i is nil")
			return
		}
	}
	fmt.Println(v.Kind())
	fmt.Println("i isn't nil")
}

func main() {
	var i string
	IsNull(i)
}
