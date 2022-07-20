package p00706

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuestion(t *testing.T) {
	ast := assert.New(t)

	obj := Constructor()
	obj.Put(1, 1)
	obj.Put(2, 2)
	ast.Equal(1, obj.Get(1))
	ast.Equal(-1, obj.Get(3))
	obj.Put(2, 1)
	ast.Equal(1, obj.Get(2))
	obj.Remove(2)
	ast.Equal(-1, obj.Get(2))
}
