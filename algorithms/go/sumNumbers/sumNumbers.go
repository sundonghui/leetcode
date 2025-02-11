package sumnumbers

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	list := []*TreeNode{root}
	leaves := []int{}
	for len(list) > 0 {
		node := list[0]
		list = list[1:]
		if node.Left == nil && node.Right == nil {
			leaves = append(leaves, node.Val)
		}
		if node.Left != nil {
			node.Left.Val += node.Val * 10
			list = append(list, node.Left)
		}
		if node.Right != nil {
			node.Right.Val += node.Val * 10
			list = append(list, node.Right)
		}
	}
	sum := 0
	for _, leaf := range leaves {
		sum += leaf
	}
	return sum
}
