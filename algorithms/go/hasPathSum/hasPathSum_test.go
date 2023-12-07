package haspathsum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_hasPathSum(t *testing.T) {
	root := TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 11,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 13,
			},
			Right: &TreeNode{
				Val: 4,
				Right: &TreeNode{
					Val: 1,
				},
			},
		},
	}
	assert.Equal(t, true, hasPathSum(&root, 22))
}

func Test_hasPathSum1(t *testing.T) {
	root := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	assert.Equal(t, false, hasPathSum(&root, 22))
}
