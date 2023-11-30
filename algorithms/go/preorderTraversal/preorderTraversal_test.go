package preordertraversal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_preorderTraversal(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
			},
		},
	}
	expected := []int{1, 2, 3}
	assert.Equal(t, expected, PreorderTraversal(root))
}
