package gameoflife

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_gameOfLife(t *testing.T) {
	expected := [][]int{{1, 1}, {1, 1}}
	board := [][]int{{1, 1}, {1, 0}}
	gameOfLife(board)
	assert.Equal(t, expected, board)
}
