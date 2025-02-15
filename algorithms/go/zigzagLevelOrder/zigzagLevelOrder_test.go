package zigzaglevelorder

import (
	"reflect"
	"testing"
)

func Test_zigzagLevelOrder_example1(t *testing.T) {
	// input
	root := &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}
	// output
	expected := [][]int{{3}, {20, 9}, {15, 7}}
	// test
	actual := zigzagLevelOrder(root)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}

func Test_zigzagLevelOrder_example2(t *testing.T) {
	// input
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 6}}}
	// output
	expected := [][]int{{1}, {3, 2}, {4, 5, 6}}
	// test
	actual := zigzagLevelOrder(root)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}
