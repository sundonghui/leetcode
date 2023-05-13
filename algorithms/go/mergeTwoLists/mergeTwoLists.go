package mergetwolists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var list []int
	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			list = append(list, list2.Val)
			list2 = list2.Next
		} else {
			list = append(list, list1.Val)
			list1 = list1.Next
		}
	}
	if list1 == nil {
		for list2 != nil {
			list = append(list, list2.Val)
			list2 = list2.Next
		}
	}
	if list2 == nil {
		for list1 != nil {
			list = append(list, list1.Val)
			list1 = list1.Next
		}
	}

	return arrayToList(list)
}

func arrayToList(array []int) *ListNode {
	if len(array) < 1 {
		return nil
	}
	head := &ListNode{Val: array[0], Next: nil}
	curr := head
	for i := 1; i < len(array); i++ {
		curr.Next = &ListNode{Val: array[i], Next: nil}
		curr = curr.Next
	}
	return head
}
