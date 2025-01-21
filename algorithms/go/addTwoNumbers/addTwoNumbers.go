package addtwonumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	node := new(ListNode)
	head := node
	adv := 0
	for l1 != nil && l2 != nil {
		tmp := l1.Val + l2.Val + adv
		adv = 0
		if tmp >= 10 {
			adv = 1
			tmp -= 10
		}

		tmpNode := ListNode{
			Val: tmp,
		}
		node.Next = &tmpNode
		node = &tmpNode

		l1 = l1.Next
		l2 = l2.Next
	}

	var advNode *ListNode
	if l2 != nil {
		advNode = l2
	}
	if l1 != nil {
		advNode = l1
	}
	for advNode != nil {
		tmp := advNode.Val + adv
		adv = 0
		if tmp >= 10 {
			adv = 1
			tmp -= 10
		}
		tmpNode := ListNode{
			Val: tmp,
		}
		node.Next = &tmpNode
		node = &tmpNode

		advNode = advNode.Next
	}
	if adv > 0 {
		tmpNode := ListNode{
			Val: adv,
		}
		node.Next = &tmpNode
		node = &tmpNode
	}

	return head.Next
}
*/

func addTwoNumbers(l1 *ListNode, l2 *ListNode) (head *ListNode) {
	var tail *ListNode
	carry := 0
	for l1 != nil || l2 != nil {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		sum, carry = sum%10, sum/10
		if head == nil {
			head = &ListNode{Val: sum}
			tail = head
		} else {
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
	}
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return
}
