package partition

import (
	"reflect"
	"testing"
)

func Test_partition_example1(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val: 2,
						},
					},
				},
			},
		},
	}
	x := 3
	expected := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 5,
						},
					},
				},
			},
		},
	}
	actual := partition(head, x)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("partition() = %v, want %v", actual, expected)
	}
}

func Test_partition_example2(t *testing.T) {
	head := &ListNode{Val: 2, Next: &ListNode{Val: 1}}
	x := 2
	expected := &ListNode{Val: 1, Next: &ListNode{Val: 2}}
	actual := partition(head, x)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("partition() = %v, want %v", actual, expected)
	}
}
