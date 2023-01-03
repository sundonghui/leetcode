package kthlargest

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthLargest(root *TreeNode, k int) int {
	res := -1

	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Right)
		k--
		if k == 0 {
			res = node.Val
			return
		}
		dfs(node.Left)
	}
	dfs(root)
	return res
}
