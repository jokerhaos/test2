package main

import (
	"fmt"
	"reflect"
)

type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Score float32
	Sex   string
}

func (s Monster) Print() {
	fmt.Println("----start-----")
	fmt.Println(s)
	fmt.Println("-----end-----")
}

func (s Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}

func (s Monster) Set(name string, age int, score float32, sex string) {
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
}

func TestStruct(b interface{}) {
	rTyp := reflect.TypeOf(b)
	fmt.Println(rTyp, rTyp.Name(), rTyp.Kind())
	rVal := reflect.ValueOf(b)
	fmt.Printf("rVal=%v rVal=%T\n", rVal, rVal)
	kd := rVal.Kind()
	if kd != reflect.Struct {
		fmt.Println("error struct")
		panic("error struct")
	}
	fmt.Println("------继续------")

	fieldNum := rVal.NumField()
	fmt.Printf("总共有多少字段：%v\n", fieldNum)
	for i := 0; i < fieldNum; i++ {
		tagVal := rTyp.Field(i).Tag.Get("json")
		fmt.Printf("Field %d , tag = %v , value：%v \n", i, tagVal, rVal.Field(i))
		if tagVal != "" {
		}
	}

	methodNum := rVal.NumMethod()
	fmt.Printf("总共有多少方法：%v\n", methodNum)

	// 调用方法会根据ASCII码排序的
	rVal.Method(1).Call(nil)

	var params []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(20))
	res := rVal.Method(0).Call(params)
	fmt.Println(res[0].Int())

}

func main() {
	a := Monster{
		Name:  "joker",
		Age:   18,
		Score: 10.2,
		Sex:   "男",
	}
	TestStruct(a)
}
