package numislands

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_numIslands_example1(t *testing.T) {
	// Input: grid = [
	//   ["1","1","1","1","0"],
	//   ["1","1","0","1","0"],
	//   ["1","1","0","0","0"],
	//   ["0","0","0","0","0"]
	// ]
	// Output: 1
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	assert.Equal(t, 1, numIslands(grid))
}

func Test_numIslands_example2(t *testing.T) {
	// Input: grid = [
	//   ["1","1","0","0","0"],
	//   ["1","1","0","0","0"],
	//   ["0","0","1","0","0"],
	//   ["0","0","0","1","1"]
	// ]
	// Output: 3
	grid := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	assert.Equal(t, 3, numIslands(grid))
}
