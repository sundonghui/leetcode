package maxdepth

import "testing"

func Test_maxDepth_example1(t *testing.T) {
	// Input: root = [3,9,20,null,null,15,7]
	// Output: 3
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	expected := 3
	if got := maxDepth(root); got != expected {
		t.Errorf("maxDepth() = %v, want %v", got, expected)
	}
}

func Test_maxDepth_example2(t *testing.T) {
	// Input: root = [1,null,2]
	// Output: 2
	root := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
		},
	}
	expected := 2
	if got := maxDepth(root); got != expected {

		t.Errorf("maxDepth() = %v, want %v", got, expected)
	}
}
