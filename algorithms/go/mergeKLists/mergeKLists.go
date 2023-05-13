package mergeklists

import (
	"sort"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	list := make([]int, 0, len(lists))
	for len(lists) > 0 {
		sort.Slice(lists, func(i, j int) bool {
			if lists[i] != nil && lists[j] != nil {
				return lists[i].Val < lists[j].Val
			}
			return true
		})
		if lists[0] != nil {
			list = append(list, lists[0].Val)
			lists[0] = lists[0].Next
		}
		if lists[0] == nil {
			lists = lists[1:]
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
