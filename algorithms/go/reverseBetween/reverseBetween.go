package reversebetween

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil {
		return nil
	}
	resultList := []int{}
	for i := 0; i < left-1; i++ {
		resultList = append(resultList, head.Val)
		head = head.Next
	}
	betweenList := []int{}
	for i := 0; i <= right-left; i++ {
		betweenList = append([]int{head.Val}, betweenList...)
		head = head.Next
	}
	resultList = append(resultList, betweenList...)
	for head != nil {
		resultList = append(resultList, head.Val)
		head = head.Next
	}
	if len(resultList) == 0 {
		return nil
	}
	result := &ListNode{Val: resultList[0]}
	current := result
	for i := 1; i < len(resultList); i++ {
		current.Next = &ListNode{Val: resultList[i]}
		current = current.Next
	}
	return result
}
