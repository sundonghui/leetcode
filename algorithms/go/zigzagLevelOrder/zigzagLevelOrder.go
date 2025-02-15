package zigzaglevelorder

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	result := [][]int{}
	list := []*TreeNode{root}
	arrow := 1
	for len(list) > 0 {
		tempList := []*TreeNode{}
		ints := []int{}
		for index, node := range list {
			if arrow > 0 {
				ints = append(ints, node.Val)
			} else {
				ints = append(ints, list[len(list)-1-index].Val)
			}
			if node.Left != nil {
				tempList = append(tempList, node.Left)
			}
			if node.Right != nil {
				tempList = append(tempList, node.Right)
			}
		}
		result = append(result, ints)
		list = tempList
		arrow *= -1
	}
	return result
}
