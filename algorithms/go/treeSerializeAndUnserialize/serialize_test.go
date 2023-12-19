package treeserializeandunserialize

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_serialize(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	ser := Constructor()
	deser := Constructor()
	data := ser.serialize(root)
	ans := deser.deserialize(data)
	assert.Equal(t, root, ans)
}
