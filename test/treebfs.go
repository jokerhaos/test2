package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Test struct {
	b string
	a int
	d string
	c float64
}

// GDAFEMHZ

// 			g
// 		d			m
// 	a		f	h		z
// 		e

func main() {
	// 构造二叉树
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}

	// 输出当前层的节点值
	fmt.Println("Current Level:")
	outputCurrentLevel(root)
}

func outputCurrentLevel(root *TreeNode) {
	if root == nil {
		return
	}

	queue := []*TreeNode{root} // 使用队列辅助进行广度优先搜索
	fmt.Printf("%#v --> %d \r\n", queue, len(queue))
	for len(queue) > 0 {
		// 记录当前层的节点数量
		levelSize := len(queue)

		// 输出当前层的节点值
		for i := 0; i < levelSize; i++ {
			node := queue[i]
			fmt.Printf("%d ", node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		fmt.Println() // 换行输出

		queue = queue[levelSize:] // 将当前层的节点出队
	}
}

func c(root *TreeNode) {
	queque := []*TreeNode{root}

	for len(queque) > 0 {
		// 记录当前层的节点数量
		size := len(queque)
		for i := 0; i < size; i++ {
			node := queque[i]
			fmt.Printf("%d ", node.Val)
			if node.Left != nil {
				queque = append(queque, node.Left)
			}
			if node.Right != nil {
				queque = append(queque, node.Right)
			}

			fmt.Println() // 换行输出
			// 将当前层的节点出队
			queque = queque[size:]
		}
	}
}
