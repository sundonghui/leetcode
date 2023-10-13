package main

import (
	"fmt"
	"math/bits"
	"math/rand"
)

func findTop100(data []int) ([]int, int) {
	comparisons := 0

	compare := func(i, j int) bool {
		comparisons++
		return data[i] > data[j]
	}

	exchange := func(i, j int) {
		data[i], data[j] = data[j], data[i]
	}

	n := len(data)
	treeSize := 1 << uint(bits.Len(uint(n))) // 计算二叉树的大小
	offset := treeSize - 1                   // 存储比较结果的偏移量

	// 构建二叉树并保存比较结果
	for i := 0; i < n; i++ {
		index := i + offset
		for index > 0 {
			parent := (index - 1) / 2
			if compare(index, parent) {
				exchange(index, parent)
			}
			index = parent
		}
	}

	top100 := []int{data[offset]} // 存储前100大的数字
	for i := 1; i < 100; i++ {
		index := 0
		for {
			leftChild := 2*index + 1
			rightChild := 2*index + 2

			if leftChild >= treeSize {
				break
			}

			if rightChild >= treeSize || compare(leftChild, rightChild) {
				if compare(leftChild, offset) {
					exchange(leftChild, offset)
				}
				index = leftChild
			} else {
				if compare(rightChild, offset) {
					exchange(rightChild, offset)
				}
				index = rightChild
			}
		}

		top100 = append(top100, data[offset])
		offset++
	}

	return top100, comparisons
}

func main() {
	data := rand.Perm(100000) // 生成1～10万的随机数
	top100, comparisons := findTop100(data)
	fmt.Println("前100大的数字：", top100)
	fmt.Println("比较的次数：", comparisons)
}
