package translatenum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTranslateNum(t *testing.T) {
	n := 12258
	expected := 5
	assert.Equal(t, expected, translateNum(n))
}
