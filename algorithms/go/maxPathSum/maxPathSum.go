package maxpathsum

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var maxPath func(node *TreeNode) int
	maxSum := math.MinInt32
	maxPath = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		maxLeft := max(0, maxPath(node.Left))
		maxRight := max(0, maxPath(node.Right))

		maxSum = max(maxSum, node.Val+maxLeft+maxRight)

		return node.Val + max(maxLeft, maxRight)
	}
	maxPath(root)
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
