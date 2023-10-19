package pathsum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPathSum(t *testing.T) {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	assert.Equal(t, [][]int{}, pathSum(tree, 5))
}

func pathSumTest(tree *TreeNode, target int) [][]int {
	paths := []int{}
	ans := [][]int{}
	var dfs func(*TreeNode, int)
	dfs = func(tn *TreeNode, i int) {
		if tn == nil {
			return
		}
		i -= tn.Val
		paths = append(paths, tn.Val)
		defer func() {
			paths = paths[:len(paths)-1]
		}()
		if tn.Left == nil && tn.Right == nil && i == 0 {
			ans = append(ans, append([]int(nil), paths...))
			return
		}
		dfs(tn.Left, i)
		dfs(tn.Right, i)
	}
	dfs(tree, target)
	return ans
}
