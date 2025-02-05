package issametree

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 100. Same Tree
// https://leetcode.com/problems/same-tree/
/*
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil && q != nil {
		return false
	}
	if p != nil && q == nil {
		return false
	}
	if p.Val == q.Val {
		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	}
	return false
}
*/

func isSameTree(p *TreeNode, q *TreeNode) bool {
	pQ := tree2List(p)
	qQ := tree2List(q)
	if len(pQ) != len(qQ) {
		return false
	}
	for i := 0; i < len(pQ); i++ {
		if pQ[i] != qQ[i] {
			return false
		}
	}
	return true
}

func tree2List(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	queue := []*TreeNode{root}
	list := []int{}
	for len(queue) > 0 {
		sz := len(queue)
		for sz > 0 {
			node := queue[0]
			queue = queue[1:]
			sz--
			if node == nil {
				list = append(list, math.MinInt)
				continue
			}
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
			list = append(list, node.Val)

		}
	}
	return list
}
