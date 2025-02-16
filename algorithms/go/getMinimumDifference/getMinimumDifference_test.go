package getminimumdifference

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getMinimumDifference_example1(t *testing.T) {
	// Input: root = [4,2,6,1,3]
	// Output: 1
	// Explanation:
	// The minimum difference is 1, which is the difference between 2 and 1 (or between 2 and 3).
	root := &TreeNode{Val: 4, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 6}}
	assert.Equal(t, 1, getMinimumDifference(root))
}

func Test_getMinimumDifference_example2(t *testing.T) {
	// Input: root = [1,0,48,null,null,12,49]
	// Output: 1
	// Explanation:
	// The minimum difference is 1, which is the difference between 1 and 0 (or between 1 and 2).
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 48, Left: &TreeNode{Val: 12}, Right: &TreeNode{Val: 49}}}
	assert.Equal(t, 1, getMinimumDifference(root))
}
