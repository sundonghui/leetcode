package sumnumbers

import "testing"

func Test_sumNumbers_example1(t *testing.T) {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	if sumNumbers(root) != 25 {
		t.Errorf("sumNumbers() = %v, want %v", sumNumbers(root), 25)
	}
}

func Test_sumNumbers_example2(t *testing.T) {
	root := &TreeNode{Val: 4, Left: &TreeNode{Val: 9, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 1}}, Right: &TreeNode{Val: 0}}
	if sumNumbers(root) != 1026 {
		t.Errorf("sumNumbers() = %v, want %v", sumNumbers(root), 1026)
	}
}
