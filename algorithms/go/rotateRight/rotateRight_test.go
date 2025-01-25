package rotateright

import (
	"reflect"
	"testing"
)

func Test_rotateRight_example1(t *testing.T) {
	// Input: head = [1,2,3,4,5], k = 2
	// Output: [4,5,1,2,3]
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}
	k := 2
	expected := &ListNode{Val: 4}
	expected.Next = &ListNode{Val: 5}
	expected.Next.Next = &ListNode{Val: 1}
	expected.Next.Next.Next = &ListNode{Val: 2}
	expected.Next.Next.Next.Next = &ListNode{Val: 3}
	actual := rotateRight(head, k)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("rotateRight() = %v, want %v", actual, expected)
	}
}

func Test_rotateRight_example2(t *testing.T) {
	// Input: head = [0,1,2], k = 4
	// Output: [2,0,1]
	head := &ListNode{Val: 0}
	head.Next = &ListNode{Val: 1}
	head.Next.Next = &ListNode{Val: 2}
	k := 4
	expected := &ListNode{Val: 2}
	expected.Next = &ListNode{Val: 0}
	expected.Next.Next = &ListNode{Val: 1}
	actual := rotateRight(head, k)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("rotateRight() = %v, want %v", actual, expected)
	}
}
