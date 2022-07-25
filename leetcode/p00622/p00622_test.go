package p00622

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuestion1(t *testing.T) {
	ast := assert.New(t)

	obj := Constructor(3)
	ast.Equal(true, obj.EnQueue(1))
	ast.Equal(true, obj.EnQueue(2))
	ast.Equal(true, obj.EnQueue(3))
	ast.Equal(false, obj.EnQueue(4))
	ast.Equal(3, obj.Rear())
	ast.Equal(true, obj.IsFull())
	ast.Equal(true, obj.DeQueue())
	ast.Equal(true, obj.EnQueue(4))
	ast.Equal(4, obj.Rear())
}

func TestQuestion2(t *testing.T) {
	ast := assert.New(t)

	obj := Constructor(3)
	ast.Equal(true, obj.EnQueue(7))
	ast.Equal(true, obj.DeQueue())
	ast.Equal(-1, obj.Front())
	ast.Equal(false, obj.DeQueue())
	ast.Equal(-1, obj.Front())
	ast.Equal(-1, obj.Rear())
	ast.Equal(true, obj.EnQueue(0))
	ast.Equal(false, obj.IsFull())
	ast.Equal(true, obj.DeQueue())
	ast.Equal(-1, obj.Rear())
	ast.Equal(true, obj.EnQueue(3))
}
