package main

import (
	"fmt"
	"unsafe"
)

type Node struct {
	Data int
	Next *Node
}
type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}

func main() {

	head := new(Node)
	head.Data = 1
	node2 := new(Node)
	node2.Data = 2
	node3 := new(Node)
	node3.Data = 3
	node4 := new(Node)
	node4.Data = 4

	head.Next = node2
	node2.Next = node3
	node3.Next = node4

	fmt.Println("翻转前：")
	printNum(head)

	result := reverse(head)
	fmt.Println("翻转后：")
	printNum(result)

}

// 反转
func reverse(currentNode *Node) *Node {
	if currentNode == nil {
		return nil
	}
	var preNode *Node
	for currentNode != nil {
		temp := currentNode.Next   //  4
		currentNode.Next = preNode // nil
		preNode = currentNode
		currentNode = temp // 4
	}
	return preNode
}

func printNum(currentNode *Node) {
	if currentNode != nil {
		fmt.Print(currentNode.Data)
		if currentNode.Next != nil {
			fmt.Print(",")
			printNum(currentNode.Next)
		} else {
			fmt.Println()
		}
	}
}
