package evalrpn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvalRPN_example1(t *testing.T) {
	tokens := []string{"2","1","+","3","*"}
	// Output: 9
	// Explanation: ((2 + 1) * 3) = 9
	assert.Equal(t, 9, evalRPN(tokens))
}

func TestEvalRPN_example2(t *testing.T) {
	tokens := []string{"4","13","5","/","+"}
	// Output: 6
	// Explanation: (4 + (13 / 5)) = 6
	assert.Equal(t, 6, evalRPN(tokens))
}

func TestEvalRPN_example3(t *testing.T) {
	tokens := []string{"10","6","9","3","+","-11","*","/","*","17","+","5","+"}
	// Output: 22
	// Explanation: ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
	assert.Equal(t, 22, evalRPN(tokens))
}