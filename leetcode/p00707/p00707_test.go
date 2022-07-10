package p00707

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuestion(t *testing.T) {
	ast := assert.New(t)

	dll := Constructor()
	dll.AddAtHead(1)
	dll.AddAtTail(3)
	dll.AddAtIndex(1, 2)
	ast.Equal(2, dll.Get(1))
	dll.DeleteAtIndex(1)
	ast.Equal(3, dll.Get(1))

	dll = Constructor()
	dll.AddAtHead(1)
	dll.AddAtTail(3)
	dll.AddAtIndex(1, 2)
	ast.Equal(2, dll.Get(1))
	dll.DeleteAtIndex(0)
	ast.Equal(2, dll.Get(0))
}
