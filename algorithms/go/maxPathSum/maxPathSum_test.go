package maxpathsum

import "testing"

func Test_maxPathSum(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	if maxPathSum(root) != 6 {
		t.Error("Test case 0 failed")
	}
}

func Test_maxPathSum1(t *testing.T) {
	root := &TreeNode{
		Val: -10,
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
	if maxPathSum(root) != 42 {
		t.Error("Test case 1 failed")
	}
}