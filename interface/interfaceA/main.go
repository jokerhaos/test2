package main

import "fmt"

type AInterface interface {
	Say()
}

type Stu struct {
}

func (stu Stu) Say() {
	fmt.Println("say")
}

type integer int

func (i integer) Say() {
	fmt.Println("integer say")
}

func main() {
	var stu Stu
	var a AInterface = stu
	a.Say()

	var i integer = 30
	var b AInterface = i
	b.Say()

}
