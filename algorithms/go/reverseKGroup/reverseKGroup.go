package reversekgroup

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	list := make([]int, 0, k)

	tmp := make([]int, 0, k)
	for head != nil {
		if len(tmp) < k {
			tmp = append(tmp, head.Val)
		} else {
			newTmp := make([]int, 0, k)
			for i := len(tmp); i > 0; i-- {
				newTmp = append(newTmp, tmp[i-1])
			}
			list = append(list, newTmp...)
			tmp = make([]int, 0, k)
			tmp = append(tmp, head.Val)
		}
		head = head.Next
	}
	if len(tmp) == k {
		for i := len(tmp); i > 0; i-- {
			list = append(list, tmp[i-1])
		}
	} else {
		list = append(list, tmp...)
	}
	return arrayToList(list)
}

func arrayToList(arr []int) *ListNode {
	if len(arr) < 1 {
		return nil
	}
	head := &ListNode{
		Val:  arr[0],
		Next: nil,
	}
	cur := head
	for i := 1; i < len(arr); i++ {
		cur.Next = &ListNode{Val: arr[i], Next: nil}
		cur = cur.Next
	}
	return head
}
