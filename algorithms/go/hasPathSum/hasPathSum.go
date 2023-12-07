package haspathsum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	return check(root, targetSum)
}

func check(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if targetSum-root.Val == 0 && root.Left == nil && root.Right == nil {
		return true
	}
	return check(root.Left, targetSum-root.Val) || check(root.Right, targetSum-root.Val)
}
