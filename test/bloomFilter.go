package main

import (
	"fmt"
	"hash/fnv"
)

// BloomFilter 是布隆过滤器的数据结构
type BloomFilter struct {
	bitArray []bool // 位数组，用于存储数据的标记位
	numHash  int    // 哈希函数的数量
}

// NewBloomFilter 创建一个新的布隆过滤器，指定大小和哈希函数数量
func NewBloomFilter(size int, numHash int) *BloomFilter {
	return &BloomFilter{
		bitArray: make([]bool, size),
		numHash:  numHash,
	}
}

// Add 将元素添加到布隆过滤器
func (bf *BloomFilter) Add(item string) {
	for i := 0; i < bf.numHash; i++ {
		index := bf.hash(item, i) % len(bf.bitArray)
		bf.bitArray[index] = true
	}
}

// Contains 检查元素是否可能在布隆过滤器中
func (bf *BloomFilter) Contains(item string) bool {
	for i := 0; i < bf.numHash; i++ {
		index := bf.hash(item, i) % len(bf.bitArray)
		if !bf.bitArray[index] {
			return false
		}
	}
	return true
}

// hash 应用哈希函数对元素进行哈希
func (bf *BloomFilter) hash(item string, seed int) uint64 {
	h := fnv.New64a()
	h.Write([]byte(item))
	h.Write([]byte(fmt.Sprint(seed)))
	return h.Sum64()
}

func main() {
	// 创建一个大小为10，使用3个哈希函数的布隆过滤器
	bf := NewBloomFilter(10, 3)

	// 向布隆过滤器中添加一些元素
	bf.Add("苹果")
	bf.Add("香蕉")
	bf.Add("橙子")

	// 检查一些元素是否可能在布隆过滤器中
	fmt.Println(bf.Contains("苹果")) // true
	fmt.Println(bf.Contains("葡萄")) // false
	fmt.Println(bf.Contains("橙子")) // true
	fmt.Println(bf.Contains("樱桃")) // false
	fmt.Println(bf.Contains("香蕉")) // true
	fmt.Println(bf.Contains("西瓜")) // false
}
