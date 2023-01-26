package lowestcommonancestor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLowestCommonAncestor(t *testing.T) {
	root := TreeNode{
		Val: 6,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 3,
				},
				Right: &TreeNode{
					Val: 5,
				},
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 7,
			},
			Right: &TreeNode{
				Val: 9,
			},
		},
	}
	r := lowestCommonAncestor(&root, &TreeNode{Val: 2}, &TreeNode{Val: 4})
	assert.Equal(t, 2, r.Val)
}
