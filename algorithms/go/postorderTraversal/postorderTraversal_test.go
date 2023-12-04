package postordertraversal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_postorderTraversal(t *testing.T) {
	root := TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}
	expected := []int{3, 2, 1}
	assert.Equal(t, expected, postorderTraversal(&root))
}
