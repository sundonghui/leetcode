package copyrandomlist

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

var cacheNode map[*Node]*Node

func deepCopy(node *Node) *Node {
	if node == nil {
		return nil
	}
	if n, has := cacheNode[node]; has {
		return n
	}
	newNode := &Node{Val: node.Val}
	cacheNode[node] = newNode
	newNode.Next = deepCopy(node.Next)
	newNode.Random = deepCopy(node.Random)
	return newNode
}

func copyRandomList(head *Node) *Node {
	cacheNode = map[*Node]*Node{}
	return deepCopy(head)
}

func CopyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	for node := head; node != nil; node = node.Next.Next {
		node.Next = &Node{Val: node.Val, Next: node.Next}
	}
	for node := head; node != nil; node = node.Next.Next {
		if node.Random != nil {
			node.Next.Random = node.Random.Next
		}
	}
	headNew := head.Next
	for node := head; node != nil; node = node.Next {
		nodeNew := node.Next
		node.Next = node.Next.Next
		if nodeNew.Next != nil {
			nodeNew.Next = nodeNew.Next.Next
		}
	}
	return headNew
}
