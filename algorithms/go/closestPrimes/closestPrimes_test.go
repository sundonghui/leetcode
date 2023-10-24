package closestprimes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_closestPrimes(t *testing.T) {
	expected := []int{11, 13}
	assert.Equal(t, expected, closestPrimes(10, 19))
}

func Test_closestPrimes1(t *testing.T) {
	expected := []int{29, 31}
	assert.Equal(t, expected, closestPrimes(19, 31))
}

func Test_closestPrimes2(t *testing.T) {
	expected := []int{29, 31}
	assert.Equal(t, expected, closestPrimes(18, 72))
}

func Test_closestPrimes3(t *testing.T) {
	expected := []int{2, 3}
	assert.Equal(t, expected, closestPrimes(1, 1000))
}

func Test_closestPrimes4(t *testing.T) {
	expected := []int{69371, 69379}
	assert.Equal(t, expected, closestPrimes(69371, 69379))
}
