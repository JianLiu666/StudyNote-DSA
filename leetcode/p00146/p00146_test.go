package p00146

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuestion(t *testing.T) {
	ast := assert.New(t)

	obj := Constructor(2)
	obj.Put(2, 1)
	obj.Put(2, 2)
	ast.Equal(2, obj.Get(2))
	obj.Put(1, 1)
	obj.Put(4, 1)
	ast.Equal(-1, obj.Get(2))
}
