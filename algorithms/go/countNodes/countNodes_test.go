package countnodes

import "testing"

func Test_countNodes_example1(t *testing.T) {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3}}
	got := countNodes(root)
	want := 5
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func Test_countNodes_example2(t *testing.T) {
	var root *TreeNode
	got := countNodes(root)
	want := 0
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func Test_countNodes_example3(t *testing.T) {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 6}}}
	got := countNodes(root)
	want := 6
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
