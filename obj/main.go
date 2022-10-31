package main

import (
	"encoding/json"
	"fmt"
)

type Cat struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Color string `json:"color"`
}

func main() {
	p1 := Cat{
		Name:  "小猫",
		Age:   1,
		Color: "白色",
	}
	fmt.Println(p1)

	// 指针，标准写法(*p2).Name = "小狗"，但是golang底层做了优化处理 p2.Name === (*p2).Name
	var p2 *Cat = new(Cat)
	p2.Name = "小狗"
	p2.Age = 2
	p2.Color = "黑色"
	fmt.Println(*p2)

	// 转json
	str, _ := json.Marshal(p1)
	fmt.Println(string(str))

	p1.test()
	p2.test()

	// 指针方法编译器底层做了优化 (&p1).update() 等价于 p1.update()
	p1.update()
	// (&p1).update();
	fmt.Println(p1)
}

func (p Cat) test() {
	fmt.Println("调用方法test", p.Name)
}

func (p *Cat) update() {
	p.Age += 1
	fmt.Println("update", *p)
}
