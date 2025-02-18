package clonegraph

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	visited := make(map[*Node]*Node)
	var clone func(node *Node) *Node
	clone = func(node *Node) *Node {
		if node == nil {
			return nil
		}
		if _, ok := visited[node]; ok {
			return visited[node]
		}
		newNode := &Node{Val: node.Val}
		visited[node] = newNode
		for _, neighbor := range node.Neighbors {
			newNode.Neighbors = append(newNode.Neighbors, clone(neighbor))
		}
		return newNode
	}
	return clone(node)
}
