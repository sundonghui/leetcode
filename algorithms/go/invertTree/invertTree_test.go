package inverttree

import (
	"reflect"
	"testing"
)

func Test_invertTree_example1(t *testing.T) {
	root := &TreeNode{Val: 4, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 7, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 9}}}
	// 4
	// /   \
	// 2     7
	// / \   / \
	// 1  3 6   9
	invertTree(root)
	expected := &TreeNode{Val: 4, Left: &TreeNode{Val: 7, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 6}}, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 1}}}
	// 4
	// /   \
	// 7     2
	// / \   / \
	// 9  6 3   1
	if !reflect.DeepEqual(root, expected) {
		t.Errorf("invertTree() = %v, want %v", root, expected)
	}
}

func Test_invertTree_example2(t *testing.T) {
	root := &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}
	// 2
	// / \
	// 1   3
	invertTree(root)
	expected := &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 1}}
	// 2
	// / \
	// 3   1
	if !reflect.DeepEqual(root, expected) {
		t.Errorf("invertTree() = %v, want %v", root, expected)
	}
}

func Test_invertTree_example3(t *testing.T) {
	root := &TreeNode{Val: 1}
	// 1
	invertTree(root)
	expected := &TreeNode{Val: 1}
	// 1
	if !reflect.DeepEqual(root, expected) {
		t.Errorf("invertTree() = %v, want %v", root, expected)
	}
}
func Test_invertTree_example4(t *testing.T) {
	root := &TreeNode{Val: 0}
	// 0
	invertTree(root)
	expected := &TreeNode{Val: 0}
	// 0
	if !reflect.DeepEqual(root, expected) {
		t.Errorf("invertTree() = %v, want %v", root, expected)
	}
}
