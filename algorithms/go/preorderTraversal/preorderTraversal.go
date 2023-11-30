package preordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	var vals []int
	var preorder func(*TreeNode)
	preorder = func(tn *TreeNode) {
		if tn == nil {
			return
		}
		vals = append(vals, tn.Val)
		preorder(tn.Left)
		preorder(tn.Right)
	}
	preorder(root)
	return vals
}

func PreorderTraversal(root *TreeNode) (vals []int) {
	stack := []*TreeNode{}
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			vals = append(vals, node.Val)
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return
}
