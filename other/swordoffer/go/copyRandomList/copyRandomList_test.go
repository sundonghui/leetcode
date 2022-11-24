package copyrandomlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCopyRandomList(t *testing.T) {
	node := &Node{
		Val: 7,
		Next: &Node{
			Val: 13,
			Next: &Node{
				Val: 11,
				Next: &Node{
					Val: 10,
					Next: &Node{
						Val: 1,
					},
					Random: &Node{
						Val: 11,
						Next: &Node{
							Val: 10,
							Next: &Node{
								Val: 11,
							},
						},
						Random: &Node{
							Val: 1,
						},
					},
				},
				Random: &Node{
					Val: 1,
					Random: &Node{
						Val: 7,
					},
				},
			},
			Random: &Node{
				Val: 7,
			},
		},
	}
	assert.Equal(t, node, copyRandomList(node))
}
