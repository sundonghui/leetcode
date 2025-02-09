package connect

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	q := []*Node{root}
	for len(q) > 0 {
		temp := q
		q = nil
		for i, node := range temp {
			if i+1 < len(temp) {
				temp[i].Next = temp[i+1]
			}
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
	}
	return root
}
