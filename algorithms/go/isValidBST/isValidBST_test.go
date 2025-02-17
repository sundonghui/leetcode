package isvalidbst

import "testing"

func Test_isValidBST_example1(t *testing.T) {
	// Input: [2,1,3]
	// Output: true
	root := &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}
	if !isValidBST(root) {
		t.Errorf("isValidBST() = %v, want %v", isValidBST(root), true)
	}
}

func Test_isValidBST_example2(t *testing.T) {
	// Input: [5,1,4,null,null,3,6]
	// Output: false
	root := &TreeNode{Val: 5, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 6}}}
	if isValidBST(root) {
		t.Errorf("isValidBST() = %v, want %v", isValidBST(root), false)
	}
}

func Test_isValidBST_example3(t *testing.T) {
	// Input: [2,2,2]
	// Output: false
	root := &TreeNode{Val: 2, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 2}}
	if isValidBST(root) {
		t.Errorf("isValidBST() = %v, want %v", isValidBST(root), false)
	}
}
