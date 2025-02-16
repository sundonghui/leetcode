package kthsmallest

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	stack := []*TreeNode{}
	for {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		stack, root = stack[:len(stack)-1], stack[len(stack)-1]
		if len(stack) == 0 {
			return 0
		}
		k--
		if k == 0 {
			return root.Val
		}
		root = root.Right
	}
}
