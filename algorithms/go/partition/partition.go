package partition

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	list := []int{}
	for head != nil {
		list = append(list, head.Val)
		head = head.Next
	}
	minList := []int{}
	maxListMap := map[int]bool{}
	for i := 0; i < len(list); i++ {
		if list[i] < x {
			minList = append(minList, list[i])
		} else {
			maxListMap[i] = true
		}
	}
	maxList := []int{}
	for i := 0; i < len(list); i++ {
		if maxListMap[i] {
			maxList = append(maxList, list[i])
		}
	}
	minList = append(minList, maxList...)
	return listToNode(minList)
}

func listToNode(list []int) *ListNode {
	if len(list) == 0 {
		return nil
	}
	head := &ListNode{Val: list[0]}
	cur := head
	for i := 1; i < len(list); i++ {
		cur.Next = &ListNode{Val: list[i]}
		cur = cur.Next
	}
	return head
}
