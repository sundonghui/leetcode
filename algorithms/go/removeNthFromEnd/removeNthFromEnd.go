package removenthfromend

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	list := []int{}
	for head != nil {
		list = append(list, head.Val)
		head = head.Next
	}
	list = append(list[:len(list)-n], list[len(list)-n+1:]...)
	if len(list) <= 0 {
		return nil
	}
	node := &ListNode{Val: list[0]}
	cur := node
	for i := 1; i < len(list); i++ {
		cur.Next = &ListNode{Val: list[i]}
		cur = cur.Next
	}
	return node
}
