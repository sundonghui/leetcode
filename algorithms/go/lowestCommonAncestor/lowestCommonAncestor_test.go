package lowestcommonancestor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lowestCommonAncestort(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
	}
	assert.Equal(t, &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
	}, lowestCommonAncestor(root, &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
	}, &TreeNode{
		Val: 2,
	}))
}
