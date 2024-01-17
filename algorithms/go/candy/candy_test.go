package candy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_candy(t *testing.T) {
	ratings := []int{1, 0, 2}
	assert.Equal(t, 5, candy(ratings))
}

func Test_candy_1(t *testing.T) {
	ratings := []int{1, 2, 2}
	assert.Equal(t, 4, candy(ratings))
}
