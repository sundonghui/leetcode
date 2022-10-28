package reverseprint

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	var list []int
	for head != nil {
		list = append([]int{head.Val}, list...)
		head = head.Next
	}
	return list
}

func reversePrint2(head *ListNode) []int {
	if head == nil {
		return nil
	}
	return append(reversePrint2(head.Next), head.Val)
}
