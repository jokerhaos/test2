package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age  int
}

type Monster struct {
	Name string
	Age  int
}

func main() {
	num := 10
	// test(&num)
	test3(&num)
	fmt.Println(num)

	// stu := Student{
	// 	Name: "tom",
	// 	Age:  18,
	// }
	// stu2 := Monster{
	// 	Name: "tom",
	// 	Age:  18,
	// }
	// test2(stu)
	// test2(stu2)
}

func test(b interface{}) {
	// 获取type,Kind
	rTyp := reflect.TypeOf(b)
	fmt.Println(rTyp, rTyp.Name(), rTyp.Kind())
	rVal := reflect.ValueOf(b)
	fmt.Println(rVal)
	fmt.Printf("rVal=%v rVal=%T\n", rVal, rVal)

	n2 := 2 + rVal.Int()
	fmt.Println(n2)

	iV := rVal.Interface()

	num2 := iV.(int)

	// n3 := num2 + n2

	// fmt.Println(n3)
	fmt.Printf("num2=%v num2=%T\n", num2, num2)
	fmt.Println(num2)

}

func test2(b interface{}) {

	// 获取type,Kind
	rTyp := reflect.TypeOf(b)
	fmt.Println(rTyp, rTyp.Name(), rTyp.Kind())
	rVal := reflect.ValueOf(b)
	fmt.Println(rVal)
	fmt.Printf("rVal=%v rVal=%T\n", rVal, rVal)
	iV := rVal.Interface()
	fmt.Printf("iV=%v iV=%T\n", iV, iV)

	// stu, Ok := iV.(Student)

	switch iV.(type) {
	case Student:
		fmt.Println("Student")

		break
	case Monster:
		fmt.Println("Monster")
		break
	}

}

func test3(b interface{}) {
	// 获取type,Kind
	rTyp := reflect.TypeOf(b)
	fmt.Println(rTyp, rTyp.Name(), rTyp.Kind())
	rVal := reflect.ValueOf(b)
	fmt.Println(rVal)
	fmt.Printf("rVal=%v rVal=%T\n", rVal, rVal)

	rVal.Elem().SetInt(20)

}

func test4(b interface{}) {

}
