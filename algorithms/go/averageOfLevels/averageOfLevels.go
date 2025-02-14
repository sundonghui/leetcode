package averageoflevels

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}
	result := []float64{}
	list := []*TreeNode{root}
	for len(list) > 0 {
		newList := []*TreeNode{}
		var total float64
		var cnt float64
		for _, node := range list {
			total += float64(node.Val)
			cnt++
			if node.Left != nil {
				newList = append(newList, node.Left)
			}
			if node.Right != nil {
				newList = append(newList, node.Right)
			}
		}
		result = append(result, total/cnt)
		list = newList
	}
	return result
}
