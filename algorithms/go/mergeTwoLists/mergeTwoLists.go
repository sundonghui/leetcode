package mergetwolists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) (head *ListNode) {
	var tail *ListNode
	for list1 != nil || list2 != nil {
		val, whichNext := getVal(list1, list2)
		if whichNext == 1 {
			list1 = list1.Next
		}
		if whichNext == 2 {
			list2 = list2.Next
		}
		if head == nil {
			head = &ListNode{Val: val}
			tail = head
		} else {
			tail.Next = &ListNode{Val: val}
			tail = tail.Next
		}
	}
	return
}

func getVal(list1 *ListNode, list2 *ListNode) (val int, whichNext int) {
	if list1 == nil {
		return list2.Val, 2
	}
	if list2 == nil {
		return list1.Val, 1
	}
	if list1.Val > list2.Val {
		return list2.Val, 2
	}
	return list1.Val, 1
}
