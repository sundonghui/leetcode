package solve

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_solve_example1(t *testing.T) {
	board := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
	}
	expected := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'X', 'X', 'X'},
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'X', 'X'},
	}
	solve(board)
	// Output: [["X","X","X","X"],["X","X","X","X"],["X","X","X","X"],["X","O","X","X"]]
	assert.Equal(t, expected, board)
}
