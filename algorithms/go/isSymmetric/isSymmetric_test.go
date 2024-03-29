package issymmetric

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isSymmetric(t *testing.T) {
	root := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
	}
	assert.Equal(t, true, isSymmetric(&root))
}
