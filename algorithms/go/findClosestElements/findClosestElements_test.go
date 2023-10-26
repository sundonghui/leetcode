package findclosestelements

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findClosestElements(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	k := 4
	x := 3
	assert.Equal(t, []int{1, 2, 3, 4}, findClosestElements(arr, k, x))
}

func Test_findClosestElements1(t *testing.T) {
	arr := []int{1, 1, 1, 10, 10, 10}
	k := 1
	x := 10
	//findClosestElements(arr,k,x)
	assert.Equal(t, []int{10}, findClosestElements(arr, k, x))
}

func Test_findClosestElements2(t *testing.T) {
	arr := []int{2, 4, 4, 4, 4, 4, 10, 15}
	k := 2
	x := 4
	//findClosestElements(arr,k,x)
	assert.Equal(t, []int{4, 4}, findClosestElements(arr, k, x))
}
