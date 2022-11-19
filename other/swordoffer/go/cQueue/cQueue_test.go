package cqueue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCQueue(t *testing.T) {	
	obj := Constructor();
	obj.AppendTail(3)
	assert.Equal(t, 3,obj.DeleteHead())
	assert.Equal(t, -1,obj.DeleteHead())
	assert.Equal(t, -1,obj.DeleteHead())
}


func TestCQueueAgain(t *testing.T) {	
	obj := Constructor();	
	assert.Equal(t, -1,obj.DeleteHead())
	obj.AppendTail(5)
	obj.AppendTail(2)
	assert.Equal(t, 5,obj.DeleteHead())
	assert.Equal(t, 2,obj.DeleteHead())
}

