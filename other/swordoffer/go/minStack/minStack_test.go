package minstack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinStack(t *testing.T)  {
	obj := Constructor();
	obj.Push(-2);
	obj.Push(0);
	obj.Push(-3);
	assert.Equal(t, -3, obj.Min())
	obj.Pop()
	assert.Equal(t,0,obj.Top())
	assert.Equal(t,-2,obj.Min())
}