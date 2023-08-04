package main

import (
	"fmt"
	"math/rand"
)

const (
	maxLevel = 16
	p        = 0.5
)

type SkipListNode struct {
	key     int
	value   int
	forward []*SkipListNode
}

type SkipList struct {
	header *SkipListNode
	level  int
}

func newNode(key, value, level int) *SkipListNode {
	return &SkipListNode{
		key:     key,
		value:   value,
		forward: make([]*SkipListNode, level, level),
	}
}

func newSkipList() *SkipList {
	header := newNode(0, 0, maxLevel)
	for i := range header.forward {
		header.forward[i] = nil
	}
	return &SkipList{
		header: header,
		level:  1,
	}
}

// 随机生成层数，根据概率 p 生成，最大不超过 maxLevel
func (sl *SkipList) randomLevel() int {
	level := 1
	for rand.Float64() < p && level < maxLevel {
		level++
	}
	return level
}

// 插入元素到跳跃表中
func (sl *SkipList) insert(key, value int) {
	update := make([]*SkipListNode, maxLevel, maxLevel)
	x := sl.header
	// 从最高层开始查找插入位置
	for i := sl.level - 1; i >= 0; i-- {
		for x.forward[i] != nil && x.forward[i].key < key {
			x = x.forward[i]
		}
		update[i] = x
	}
	level := sl.randomLevel()
	if level > sl.level {
		// 如果随机生成的层数大于当前层数，则从当前层数增长到随机生成的层数，用 header 节点填充
		for i := sl.level; i < level; i++ {
			update[i] = sl.header
		}
		sl.level = level
	}
	x = newNode(key, value, level)
	for i := 0; i < level; i++ {
		// 修改前后节点的连接关系
		x.forward[i] = update[i].forward[i]
		update[i].forward[i] = x
	}
}

// 在跳跃表中查找指定 key 的节点
func (sl *SkipList) search(key int) *SkipListNode {
	x := sl.header
	// 从最高层开始查找
	for i := sl.level - 1; i >= 0; i-- {
		for x.forward[i] != nil && x.forward[i].key < key {
			x = x.forward[i]
		}
	}
	x = x.forward[0]
	if x != nil && x.key == key {
		// 如果找到指定 key 的节点，则返回该节点
		return x
	}
	// 否则返回 nil 表示未找到
	return nil
}

// 从跳跃表中删除指定 key 的节点
func (sl *SkipList) delete(key int) {
	update := make([]*SkipListNode, maxLevel, maxLevel)
	x := sl.header
	// 从最高层开始查找需要删除的节点
	for i := sl.level - 1; i >= 0; i-- {
		for x.forward[i] != nil && x.forward[i].key < key {
			x = x.forward[i]
		}
		update[i] = x
	}
	x = x.forward[0]
	if x != nil && x.key == key {
		// 找到需要删除的节点后，修改前后节点的连接关系
		for i := 0; i < sl.level; i++ {
			if update[i].forward[i] != x {
				break
			}
			update[i].forward[i] = x.forward[i]
		}
		// 更新跳跃表的层数
		for sl.level > 1 && sl.header.forward[sl.level-1] == nil {
			sl.level--
		}
	}
}

// 打印跳跃表的内容
func (sl *SkipList) print() {
	for i := sl.level - 1; i >= 0; i-- {
		x := sl.header.forward[i]
		fmt.Printf("Level %d: ", i)
		for x != nil {
			fmt.Printf("%d ", x.key)
			x = x.forward[i]
		}
		fmt.Println()
	}
}

func main() {
	sl := newSkipList()

	for i := 1; i <= 20; i++ {
		sl.insert(i, i*10)
	}
	sl.print()

	key := 10
	node := sl.search(key)
	if node != nil {
		fmt.Printf("Key %d found, value: %d\n", key, node.value)
	} else {
		fmt.Printf("Key %d not found\n", key)
	}

	keyToDelete := 15
	sl.delete(keyToDelete)
	sl.print()
}
