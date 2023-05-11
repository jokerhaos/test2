package main

import (
	"fmt"
	"strings"
)

type Node struct {
	id       int
	parentid int
	left     int
	right    int
}

type BST struct {
	data []Node
}

func NewBST(data []Node) *BST {
	return &BST{data: data}
}

// 子级树渲染
func (t *BST) findChild(pid int) []map[string]int {
	var data []map[string]int
	for k, v := range t.data {
		if v.parentid == pid {
			m := make(map[string]int)
			m["id"] = v.id
			m["k"] = k
			data = append(data, m)
		}
	}
	return data
}

// 生成树形图
func (t *BST) toTree(pid, left int, key interface{}) int {
	right := left + 1
	childs := t.findChild(pid)
	for _, v := range childs {
		id := v["id"]
		k := v["k"]
		right = t.toTree(id, right, k)
	}
	if key != nil {
		if k, ok := key.(int); ok && k < len(t.data) {
			// 第一次进来
			t.data[k].left = left
			t.data[k].right = right
		}
	}
	return right + 1
}

func main() {
	data := []Node{
		{id: 1, parentid: 0},
		{id: 2, parentid: 1},
		{id: 3, parentid: 2},
		{id: 4, parentid: 3},
		{id: 5, parentid: 4},
		{id: 6, parentid: 3},
		{id: 7, parentid: 2},
		{id: 8, parentid: 2},
		{id: 9, parentid: 1},
		{id: 10, parentid: 1},
		{id: 11, parentid: 10},
		{id: 12, parentid: 10},
		{id: 13, parentid: 10},
		{id: 14, parentid: 10},
		{id: 15, parentid: 1},
		{id: 16, parentid: 0},
	}

	bst := NewBST(data)
	bst.toTree(0, 0, nil)
	fmt.Printf("%#v \r\n ", bst.data)
	// 输出树形结构
	for _, v := range bst.data {
		fmt.Printf("%s%d %s\n", strings.Repeat("  ", v.left/2), v.id, strings.Repeat("-", (v.right-v.left)/2))
	}
}
