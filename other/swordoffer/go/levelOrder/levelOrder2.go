package levelorder

func levelOrder2(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	q := []*TreeNode{root}
	r := [][]int{}
	for i := 0; len(q) > 0; i++ {
		r = append(r, []int{})
		p := []*TreeNode{}
		for j := 0; j < len(q); j++ {
			node :=  q[j]
			r[i] = append(r[i], node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
	}
	return r
}

func levelOrder3(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	queue := []*TreeNode{root}
	r := [][]int{}
	for i := 0; len(queue) > 0; i++ {
		vals := []int{}
        q := queue
        queue = nil
		for _, node := range q {
            vals = append(vals, node.Val)
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
		if i%2 == 1 {
            for i, n := 0, len(vals); i < n/2; i++ {
                vals[i], vals[n-1-i] = vals[n-1-i], vals[i]
            }
        }
		r = append(r, vals)
	}
	return r
}