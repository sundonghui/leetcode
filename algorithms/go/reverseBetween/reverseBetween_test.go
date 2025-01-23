package reversebetween

import "testing"

func Test_reverseBetween_example(t *testing.T) {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	result := reverseBetween(head, 2, 4)
	// Output: [1,4,3,2,5]
	if result.Val != 1 || result.Next.Val != 4 || result.Next.Next.Val != 3 || result.Next.Next.Next.Val != 2 || result.Next.Next.Next.Next.Val != 5 {
		t.Errorf("reverseBetween() = %v, want %v", head, "[1,4,3,2,5]")
	}
}
