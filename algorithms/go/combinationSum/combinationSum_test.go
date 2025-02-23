package combinationsum

import "testing"

func Test_combinationSum_example1(t *testing.T) {
	candidates := []int{2, 3, 6, 7}
	target := 7
	expected := [][]int{{2, 2, 3}, {7}}
	actual := combinationSum(candidates, target)
	if len(actual) != len(expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func Test_combinationSum_exmaple2(t *testing.T) {
	candidates := []int{2, 3, 5}
	target := 8
	expected := [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}
	actual := combinationSum(candidates, target)
	if len(actual) != len(expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func Test_combinationSum_exmaple3(t *testing.T) {
	candidates := []int{2}
	target := 1
	expected := [][]int{}
	actual := combinationSum(candidates, target)
	if len(actual) != len(expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func Test_combinationSum_exmaple4(t *testing.T) {
	candidates := []int{8, 7, 4, 3}
	target := 11
	expected := [][]int{{8, 3}, {7, 4}, {4, 4, 3}}
	actual := combinationSum(candidates, target)
	if len(actual) != len(expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}
