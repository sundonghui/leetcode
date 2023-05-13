package swappairs

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	new := head.Next
	head.Next = swapPairs(new.Next)
	new.Next = head
	return new
}
