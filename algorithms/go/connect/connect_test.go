package connect

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_connect(t *testing.T) {
	root := Node{
		Val: 1,
		Left: &Node{
			Val: 2,
			Left: &Node{
				Val: 4,
			},
			Right: &Node{
				Val: 5,
			},
		},
		Right: &Node{
			Val: 3,
			Left: &Node{
				Val: 6,
			},
			Right: &Node{
				Val: 7,
			},
		},
	}

	assert.Equal(t, connect(&root), connect(&root))
}
