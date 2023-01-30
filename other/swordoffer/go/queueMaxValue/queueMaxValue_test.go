package queuemaxvalue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxQueue(t *testing.T) {
	obj := Constructor()
	param_1 := obj.Max_value()
	assert.Equal(t, -1, param_1)

	obj.Push_back(1)
	obj.Push_back(2)
	obj.Push_back(3)
	param_3 := obj.Pop_front()
	assert.Equal(t, 1, param_3)
}
