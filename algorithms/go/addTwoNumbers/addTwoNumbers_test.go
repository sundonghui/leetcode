package addtwonumbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddTwoNumbers(t *testing.T) {
	l1 := ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
			},
		},
	}
	l2 := ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	expected := &ListNode{
		Val: 7,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val:  8,
				Next: nil,
			},
		},
	}
	r := addTwoNumbers(&l1, &l2)
	for r != nil {
		assert.Equal(t, expected.Val, r.Val)
		expected = expected.Next
		r = r.Next
	}
}

func TestAddTwoNumbers2(t *testing.T) {
	l1 := ListNode{
		Val: 0,
	}
	l2 := ListNode{
		Val: 0,
	}

	expected := &ListNode{
		Val: 0,
	}
	r := addTwoNumbers(&l1, &l2)
	for r != nil {
		assert.Equal(t, expected.Val, r.Val)
		expected = expected.Next
		r = r.Next
	}
}

func TestAddTwoNumbers3(t *testing.T) {
	l1 := ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 9,
						Next: &ListNode{
							Val: 9,
							Next: &ListNode{
								Val: 9,
							}}}}, // a list of length 6
			},
		},
	}
	l2 := ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 9,
				},
			},
		},
	}
	expected := &ListNode{
		Val: 8,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val: 0,
						Next: &ListNode{
							Val: 0,
							Next: &ListNode{
								Val:  1,
								Next: nil}}}}, // a list of length 5
			},
		},
	}
	r := addTwoNumbers(&l1, &l2)
	for r != nil {

		assert.Equal(t, expected.Val, r.Val)
		expected = expected.Next
		r = r.Next
	}
}
