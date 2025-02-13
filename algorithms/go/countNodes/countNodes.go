package countnodes

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	cnt := 0
	list := []*TreeNode{root}
	for len(list) > 0 {
		node := list[0]
		cnt++
		list = list[1:]
		if node.Left != nil {
			list = append(list, node.Left)
		}
		if node.Right != nil {
			list = append(list, node.Right)
		}
	}
	return cnt
}
