package combine

import "testing"

func Test_combine_example1(t *testing.T) {
	// Input: n = 4, k = 2
	// Output: [[1,2],[1,3],[1,4],[2,3],[2,4],[3,4]]
	// Explanation: There are 6 possible combinations of 2 numbers out of [1,2,3,4].
	n, k := 4, 2
	if len(combine(n, k)) != 6 {
		t.Errorf("combine(%d, %d) = %v, want %v", n, k, combine(n, k), 6)
	}
}

func Test_combine_example2(t *testing.T) {
	// Input: n = 1, k = 1
	// Output: [[1]]
	n, k := 1, 1
	if len(combine(n, k)) != 1 {
		t.Errorf("combine(%d, %d) = %v, want %v", n, k, combine(n, k), 1)
	}
}

func Test_combine_example3(t *testing.T) {
	// Input: n = 5, k = 3
	// Output: [[1,2,3],[1,2,4],[1,2,5],[1,3,4],[1,3,5],[1,4,5],[2,3,4],[2,3,5],[2,4,5],[3,4,5]]
	n, k := 5, 3
	if len(combine(n, k)) != 10 {
		t.Errorf("combine(%d, %d) = %v, want %v", n, k, combine(n, k), 10)
	}
}
