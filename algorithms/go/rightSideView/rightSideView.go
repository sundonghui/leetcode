package rightsideview

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	layers := [][]*TreeNode{}
	list := []*TreeNode{root}
	for len(list) > 0 {
		layers = append(layers, list)
		newList := []*TreeNode{}
		for _, node := range list {
			if node.Left != nil {
				newList = append(newList, node.Left)
			}
			if node.Right != nil {
				newList = append(newList, node.Right)
			}
		}
		list = newList
	}
	result := []int{}
	for _, layer := range layers {
		result = append(result, layer[len(layer)-1].Val)
	}
	return result
}
