package inordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node != nil {
			// 左->中->右
			inorder(node.Left)
			res = append(res, node.Val)
			inorder(node.Right)
		}
	}
	inorder(root)
	return res
}
