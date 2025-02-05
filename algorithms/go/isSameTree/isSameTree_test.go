package issametree

import "testing"

func Test_isSameTree(t *testing.T) {
	p := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	q := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	//fmt.Println(isSameTree(p, q))
	if !isSameTree(p, q) {
		t.Errorf("isSameTree() = %v, want %v", isSameTree(p, q), true)
	}
}

func Test_isSameTree2(t *testing.T) {
	p := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}
	q := &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}
	//fmt.Println(isSameTree(p, q))
	if isSameTree(p, q) {
		t.Errorf("isSameTree() = %v, want %v", isSameTree(p, q), false)
	}
}

func Test_isSameTree3(t *testing.T) {
	p := &TreeNode{Val: 1}
	q := &TreeNode{Val: 2}
	//fmt.Println(isSameTree(p, q))
	if isSameTree(p, q) {
		t.Errorf("isSameTree() = %v, want %v", isSameTree(p, q), false)
	}
}
