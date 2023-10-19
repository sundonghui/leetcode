package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_test(t *testing.T) {
	main1()
}

func Test_test1(t *testing.T) {
	main1()
}

func Test_test2(t *testing.T) {
	input := "{}{{}{}}{{{}}}{{}{{}}}"
	output := "{{{}}}"
	//main2(input)
	assert.Equal(t, output, main2(input))
}

func Test_test4(t *testing.T) {
	assert.Equal(t, 50005000, bufferChan())
}
