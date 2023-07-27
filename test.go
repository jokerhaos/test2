package main

import (
	"fmt"
	"reflect"
	"test/util"
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

type MyStruct struct{}

func (s MyStruct) Hello(name string) string {
	return "Hello, " + name + "!"
}

func (s *MyStruct) Add(a, b int) int {
	return a + b
}

func Add(a, b int) int {
	return a + b
}

func init() {
	// 注册
	util.RegisterFunction("main", "Add", Add)
}

func main() {

	// 通过包名和方法名动态调用
	packageName := "main"
	funcName := "Add"
	args := []interface{}{10, 20}
	results, err := util.Eval(packageName, funcName, args...)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	sum := results[0].(int)
	fmt.Println("Result:", sum) // Output: Result: 30

	// 通过结构体动态调用地下方法
	instance := MyStruct{}

	// 调用 Hello 方法
	result1, _ := util.CallMethod(instance, "Hello", "John")
	fmt.Println(result1[0].(string)) // Output: Hello, John!

	// 调用 Add 方法
	args2 := []interface{}{10, 20}
	result2, _ := util.CallMethod(&instance, "Add", args2...)
	fmt.Println(result2[0].(int)) // Output: 30

	// fmt.Println(strings.Split("", "/")[0])

	// a := make([]int, 0, 10)

	// fmt.Println(a)

	// b := make([]int, 10)

	// fmt.Println(b)

	// a := A{
	// 	Name:  "Alice",
	// 	Age:   intPtr(25),
	// 	Email: "",
	// 	CCC:   "ccc",
	// }

	// b := B{
	// 	Name:  "Alice2",
	// 	Email: "john@example.com",
	// 	BBB:   "",
	// }

	// CopyFields(&a, b)

	// fmt.Printf("%#v", a)
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
