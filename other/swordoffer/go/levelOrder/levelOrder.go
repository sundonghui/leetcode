package levelorder

import (
	"container/list"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
	
func levelOrder(root *TreeNode) []int {	
	for root == nil {
		return []int{}
	}

	ans := []int{}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() != 0 {
		tmp := queue.Front().Value.(*TreeNode)
		ans = append(ans, tmp.Val)
		queue.Remove(queue.Front())
		if tmp.Left != nil {
			queue.PushBack(tmp.Left)
		}
		if tmp.Right != nil {
			queue.PushBack(tmp.Right)
		}
	}
	return ans
}

func levelOrder2(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	
	ans := []int{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		q := queue
		queue = nil
		for _, tree := range q {
			ans = append(ans, tree.Val)
			if tree.Left != nil {
				queue = append(queue, tree.Left)
			}
			if tree.Right != nil {
				queue = append(queue, tree.Right)
			}
		}		
	}
	return ans
}