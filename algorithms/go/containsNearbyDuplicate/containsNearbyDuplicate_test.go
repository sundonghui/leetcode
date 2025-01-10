package containsnearbyduplicate

import "testing"

func TestContainsNearbyDuplicate(t *testing.T) {
	tests := []struct {
		nums     []int
		k        int
		expected bool
	}{
		{
			[]int{1, 2, 3, 1},
			3,
			true,
		},
		{
			[]int{1, 0, 1, 1},
			1,
			true,
		},
		{
			[]int{1, 2, 3, 1, 2, 3},
			2,
			false,
		},
	}
	for _, test := range tests {
		result := containsNearbyDuplicate(test.nums, test.k)
		if result != test.expected {
			t.Errorf("containsNearbyDuplicate(%v, %d) = %t; want %t", test.nums, test.k, result, test.expected)
		}
	}
}
