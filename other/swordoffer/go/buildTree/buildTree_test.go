package buildtree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildTree(t *testing.T) {
	preorder := []int{
		3, 9, 20, 15, 7,
	}
	inorder := []int{
		9, 3, 15, 20, 7,
	}
	expected := TreeNode{
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
	assert.Equal(t, &expected, buildTree(preorder, inorder))
}
