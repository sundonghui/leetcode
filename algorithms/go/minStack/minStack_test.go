package minstack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinStack(t *testing.T) {
	// Your MinStack object will be instantiated and called as such:
	val := 1
	obj := Constructor();
	obj.Push(val);
	obj.Pop();
	param_3 := obj.Top();
	assert.Equal(t, param_3, 0)
	obj.Push(val);
	param_4 := obj.GetMin();
	assert.Equal(t, param_4, 1)
}