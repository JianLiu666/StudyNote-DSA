package p00225

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuestion(t *testing.T) {
	ast := assert.New(t)

	myStack := Constructor()
	myStack.Push(1)
	myStack.Push(2)
	ast.Equal(2, myStack.Top())
	ast.Equal(2, myStack.Pop())
	ast.Equal(false, myStack.Empty())
}
