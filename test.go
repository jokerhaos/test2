package main

import (
	"fmt"
	"reflect"
	"time"
)

const shortDuration = 1 * time.Millisecond

type A struct {
	Name  string
	Age   *int
	Email string
	CCC   string
}

type B struct {
	Name  string
	Age   *int
	Email string
	BBB   string
}

func main() {
	a := A{
		Name:  "Alice",
		Age:   intPtr(25),
		Email: "",
		CCC:   "ccc",
	}

	b := B{
		Name:  "Alice2",
		Email: "john@example.com",
		BBB:   "",
	}

	CopyFields(&a, b)

	fmt.Printf("%#v", a)
}

func CopyFields(a interface{}, b interface{}, fields ...string) (err error) {
	at := reflect.TypeOf(a)
	av := reflect.ValueOf(a)
	bt := reflect.TypeOf(b)
	bv := reflect.ValueOf(b)
	// 简单判断下
	if at.Kind() != reflect.Ptr {
		err = fmt.Errorf("a must be a struct pointer")
		return
	}
	av = reflect.ValueOf(av.Interface())
	// 要复制哪些字段
	_fields := make([]string, 0)
	if len(fields) > 0 {
		_fields = fields
	} else {
		for i := 0; i < bv.NumField(); i++ {
			_fields = append(_fields, bt.Field(i).Name)
		}
	}
	if len(_fields) == 0 {
		fmt.Println("no fields to copy")
		return
	}
	// 复制
	for i := 0; i < len(_fields); i++ {
		name := _fields[i]
		f := av.Elem()
		if f.Kind() == reflect.Ptr {
			f = f.Elem()
		}
		f = f.FieldByName(name)
		bValue := bv.FieldByName(name)
		// a中有同名的字段并且类型一致才复制,并且b是指针类型为nil不复制
		if !(bValue.Kind() == reflect.Ptr && bValue.IsNil()) && f.IsValid() && f.Kind() == bValue.Kind() {
			f.Set(bValue)
		} else {
			fmt.Printf("no such field or different kind, fieldName: %s\n", name)
		}
	}
	return
}

func copyNonNullFields(a interface{}, b interface{}) {
	aValue := reflect.ValueOf(a).Elem()
	bValue := reflect.ValueOf(b).Elem()

	for i := 0; i < aValue.NumField(); i++ {
		aField := aValue.Field(i)
		bField := bValue.Field(i)

		// 如果 B 的字段是非空值，则将其赋值给 A
		if !bField.IsNil() && aField.IsValid() {
			bFieldValue := reflect.Indirect(bField)
			aField.Set(bFieldValue)
		}
	}
}

// 辅助函数，将字符串转换为指针
func stringPtr(s string) *string {
	return &s
}

// 辅助函数，将整数转换为指针
func intPtr(i int) *int {
	return &i
}
