package p00232

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuestion(t *testing.T) {
	ast := assert.New(t)

	myQueue := Constructor()
	myQueue.Push(1)
	myQueue.Push(2)
	ast.Equal(1, myQueue.Peek())
	ast.Equal(1, myQueue.Pop())
	ast.Equal(false, myQueue.Empty())
}
