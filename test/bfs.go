package main

import (
	"fmt"
)

// 定义图的邻接表表示法结构体
type Graph struct {
	vertices map[int][]int
}

// 添加边
func (g *Graph) AddEdge(u, v int) {
	g.vertices[u] = append(g.vertices[u], v)
}

// BFS 遍历
func (g *Graph) BFS(start int) {
	visited := make(map[int]bool) // 用于记录已访问的节点
	queue := []int{start}         // 使用队列进行层次遍历

	// 遍历队列中的节点
	for len(queue) > 0 {
		levelSize := len(queue) // 记录当前层节点数量
		fmt.Printf("levelSize => %d visited-len => %d ", levelSize, len(visited))

		// 输出当前层的节点值
		fmt.Printf("Level %d: ", len(visited)/levelSize+1)
		for i := 0; i < levelSize; i++ {
			node := queue[i]
			fmt.Printf("%d ", node)

			// 将当前节点的未访问邻居节点加入队列
			for _, neighbor := range g.vertices[node] {
				if !visited[neighbor] {
					visited[neighbor] = true
					queue = append(queue, neighbor)
				}
			}
		}
		fmt.Println("")
		// 更新队列，移除已访问的节点
		queue = queue[levelSize:]
	}
	fmt.Println()
}

func main() {
	// 初始化图
	graph := Graph{
		vertices: make(map[int][]int),
	}
	// 添加图的边
	graph.AddEdge(0, 1)
	graph.AddEdge(0, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(1, 4)
	graph.AddEdge(2, 5)
	fmt.Printf("%v \r\n", graph)
	// 使用 BFS 输出各层的值
	graph.BFS(0)
}
