package issubstructure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSubStructure(t *testing.T) {
	a := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 4,
			},
		},
	}
	b := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 4,
		},
	}

	assert.Equal(t, true, isSubStructure(a, b))
}
