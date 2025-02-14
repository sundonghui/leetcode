package averageoflevels

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_averageOfLevels_example1(t *testing.T) {
	// Input: root = [3,9,20,null,null,15,7]
	// Output: [3.00000,14.50000,11.00000]
	// Explanation: The average value of nodes on level 0 is 3, on level 1 is 14.5, and on level 2 is 11.
	root := &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}
	assert.Equal(t, []float64{3, 14.5, 11}, averageOfLevels(root))
}

func Test_averageOfLevels_example2(t *testing.T) {
	// Input: root = [3,9,20,null,null,15,7]
	// Output: [3.00000,14.50000,11.00000]
	// Explanation: The average value of nodes on level 0 is 3, on level 1 is 14.5, and on level 2 is 11.
	root := &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}
	assert.Equal(t, []float64{3, 14.5, 11}, averageOfLevels(root))
}
