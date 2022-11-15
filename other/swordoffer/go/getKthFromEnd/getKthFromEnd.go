package getkthfromend

// Definition for singly-linked list.
type ListNode struct {
	Val int
	Next *ListNode
}
 
 func getKthFromEnd(head *ListNode, k int) *ListNode {
	i := 0
	prev := head
	for head != nil {
		i++
		if i > k {
			prev = prev.Next
		}
		head = head.Next
	}
	return prev
}