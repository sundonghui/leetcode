package levelorder

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	// 为了通过测试用例 34，必须加这一行。
	if root == nil {
		return nil
	}
	var res [][]int
	list := []*TreeNode{root}
	for len(list) > 0 {
		levelList := []int{}
		tmpList := []*TreeNode{}
		for _, v := range list {
			if v != nil {
				levelList = append(levelList, v.Val)
				if v.Left != nil {
					tmpList = append(tmpList, v.Left)
				}
				if v.Right != nil {
					tmpList = append(tmpList, v.Right)
				}
			}
		}
		res = append(res, levelList)
		list = tmpList
	}
	return res
}
