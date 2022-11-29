package findmedian

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMedian(t *testing.T) {
	expected := 2.0

	obj := Constructor()
	obj.AddNum(1)
	obj.AddNum(2)
	obj.AddNum(3)
	assert.Equal(t, expected, obj.FindMedian())
}