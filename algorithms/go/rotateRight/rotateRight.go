package rotateright

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}
	list := []int{}
	for head != nil {
		list = append(list, head.Val)
		head = head.Next
	}
	total := len(list)
	newList := make([]int, len(list))
	for i, v := range list {
		if i+k < total {
			newList[i+k] = v
		} else {
			newIndex := (i + k) % total
			newList[newIndex] = v
		}
	}
	return listToListNode(newList)
}

func listToListNode(list []int) *ListNode {
	if len(list) == 0 {
		return nil
	}
	head := &ListNode{Val: list[0]}
	current := head
	for i := 1; i < len(list); i++ {
		current.Next = &ListNode{Val: list[i]}
		current = current.Next
	}
	return head
}
