package main

import (
	"fmt"

	"test/factory/model"
)

type Factory struct {
}

func main() {
	stu := model.NewStudent("joker", 12.1)

	fmt.Println(*stu)
	fmt.Println(stu.GetName())
}
