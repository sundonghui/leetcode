package postordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	res := []int{}
	var postorder func(*TreeNode)
	postorder = func(tn *TreeNode) {
		if tn == nil {
			return
		}
		postorder(tn.Left)
		postorder(tn.Right)
		res = append(res, tn.Val)
	}
	postorder(root)
	return res
}
