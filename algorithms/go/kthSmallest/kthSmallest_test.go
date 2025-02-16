package kthsmallest

import "testing"

func Test_kthSmallest_exmaple1(t *testing.T) {
	// Input: root = [3,1,4,null,2], k = 1
	// Output: 1
	root := &TreeNode{Val: 3, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 2}}}
	k := 1
	if kthSmallest(root, k) != 1 {
		t.Errorf("kthSmallest() = %v, want %v", kthSmallest(root, k), 1)
	}
}

func Test_kthSmallest_exmaple2(t *testing.T) {
	// Input: root = [5,3,6,2,4,null,null,1], k = 3
	// Output: 3
	root := &TreeNode{Val: 5, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}}, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 6}}
	k := 3
	if kthSmallest(root, k) != 3 {
		t.Errorf("kthSmallest() = %v, want %v", kthSmallest(root, k), 3)
	}
}

func Test_kthSmallest_exmaple3(t *testing.T) {
	// Input: root = [5,3,6,2,4,null,null,1], k = 10
	// Output: 0
	root := &TreeNode{Val: 5, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}}, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 6}}
	k := 10
	if kthSmallest(root, k) != 0 {
		t.Errorf("kthSmallest() = %v, want %v", kthSmallest(root, k), 0)
	}
}
