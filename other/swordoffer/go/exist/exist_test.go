package exist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExist(t *testing.T) {
	board := [][]byte{
		[]byte{'A', 'B', 'C', 'E'},
		[]byte{'S', 'F', 'C', 'S'},
		[]byte{'A', 'D', 'E', 'E'},
	}
	word := "ABCCED"

	assert.Equal(t, true, exist(board, word))
}
