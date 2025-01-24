package deleteduplicates

import (
	"reflect"
	"testing"
)

func TestDeleteDuplicates(t *testing.T) {
	// Test case 1
	head1 := &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2}}}
	expected1 := &ListNode{Val: 2}
	result1 := deleteDuplicates(head1)
	if !reflect.DeepEqual(result1, expected1) {
		t.Errorf("Test case 1 failed. Expected %v, but got %v", expected1, result1)
	}
}

func TestDeleteDuplicates2(t *testing.T) {
	// Test case 1
	head1 := &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 3}}}}}
	expected1 := &ListNode{Val: 2}
	result1 := deleteDuplicates(head1)
	if !reflect.DeepEqual(result1, expected1) {
		t.Errorf("Test case 1 failed. Expected %v, but got %v", expected1, result1)
	}
}

func TestDeleteDuplicates3(t *testing.T) {
	// Test case 1
	head1 := &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 1,
		Next: &ListNode{Val: 2,
			Next: &ListNode{Val: 3, Next: &ListNode{Val: 3,
				Next: &ListNode{Val: 4, Next: &ListNode{Val: 4, Next: &ListNode{Val: 4,
					Next: &ListNode{Val: 5}}}}}}}}}}}
	expected1 := &ListNode{Val: 2, Next: &ListNode{Val: 5}}
	result1 := deleteDuplicates(head1)
	if !reflect.DeepEqual(result1, expected1) {
		t.Errorf("Test case 1 failed. Expected %v, but got %v", expected1, result1)
	}
}
