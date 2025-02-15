package levelorder

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	list := []*TreeNode{root}
	resultList := [][]int{}
	for len(list) > 0 {
		intList := []int{}
		tempList := make([]*TreeNode, 0, len(list))
		for _, node := range list {
			intList = append(intList, node.Val)
			if node.Left != nil {
				tempList = append(tempList, node.Left)
			}
			if node.Right != nil {
				tempList = append(tempList, node.Right)
			}
		}
		list = tempList
		resultList = append(resultList, intList)
	}
	return resultList
}
