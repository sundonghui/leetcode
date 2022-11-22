package deletenode

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(head *ListNode, val int) *ListNode {
	if head.Val == val {
		return head.Next
	}
	head.Next = deleteNode(head.Next, val)
	return head
}

func deleteNode1(head *ListNode, val int) *ListNode {
	if head.Val == val {
		return head.Next
	}
	prev, cur := head, head
	for cur != nil && cur.Val != val {
		prev, cur = cur, cur.Next
	}
	if cur != nil {
		prev.Next = cur.Next
	}
	return head
}
