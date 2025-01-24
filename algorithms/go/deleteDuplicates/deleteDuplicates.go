package deleteduplicates

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	list := []int{}
	m := map[int]int{}
	for head != nil {
		list = append(list, head.Val)
		m[head.Val]++
		head = head.Next
	}
	if len(list) == 0 {
		return nil
	}
	uniqueList := make([]int, 0, len(list))
	for _, v := range list {
		if m[v] == 1 {
			uniqueList = append(uniqueList, v)
		}
	}
	if len(uniqueList) == 0 {
		return nil
	}
	return listToListNode(uniqueList)
}

func listToListNode(list []int) *ListNode {
	head := &ListNode{Val: list[0]}
	cur := head
	for i := 1; i < len(list); i++ {
		cur.Next = &ListNode{Val: list[i]}
		cur = cur.Next
	}
	return head
}
