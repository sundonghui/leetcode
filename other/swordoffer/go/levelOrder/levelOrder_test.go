package levelorder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLevelOrder(t *testing.T) {
	root := &TreeNode{
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
	expected := []int{3,9,20,15,7}
	assert.Equal(t, expected, levelOrder(root))

	root1 := &TreeNode{
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
	expected1 := []int{3,9,20,15,7}
	assert.Equal(t, expected1, LevelOrder(root1))
}


func TestLevelOrder2(t *testing.T) {
	root := &TreeNode{
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
	expected := [][]int{
		{3},
		{9, 20},
		{15, 7},
	}
	assert.Equal(t, expected, levelOrder2(root))
}

func TestLevelOrder3(t *testing.T) {
	root := &TreeNode{
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
	expected := [][]int{
		{3},
		{20, 9},
		{15, 7},
	}
	assert.Equal(t, expected, levelOrder3(root))
}