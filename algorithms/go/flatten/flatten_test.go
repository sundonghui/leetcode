package flatten

import "testing"

func Test_flatten(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 5}
	root.Left.Left = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 4}
	root.Right.Right = &TreeNode{Val: 6}
	// fmt.Println(root)
	flatten(root)
	// fmt.Println(root)
	expected := []int{1, 2, 3, 4, 5, 6}
	for _, v := range expected {
		if v != root.Val {
			t.Errorf("expected %d, got %d", v, root.Val)
		}
		root = root.Right
	}
}
