package rightsideview

import (
	"reflect"
	"testing"
)

func Test_rightSideView_example1(t *testing.T) {
	// Input: [1,2,3,null,5,null,4]
	// Output: [1,3,4]
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 4,
			},
		},
	}
	result := rightSideView(root)
	expected := []int{1, 3, 4}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func Test_rightSideView_example2(t *testing.T) {
	root := &TreeNode{Val: 1}
	result := rightSideView(root)
	expected := []int{1}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func Test_rightSideView_example3(t *testing.T) {
	root := &TreeNode{}
	result := rightSideView(root)
	expected := []int{}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func Test_rightSideView_example4(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 5,
				},
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	result := rightSideView(root)
	expected := []int{1, 3, 4, 5}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
