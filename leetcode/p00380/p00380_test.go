package p00380

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuestion(t *testing.T) {
	ast := assert.New(t)

	obj := Constructor()
	ast.Equal(true, obj.Insert(1))
	ast.Equal(false, obj.Remove(2))
	ast.Equal(true, obj.Insert(2))
	obj.GetRandom()
	ast.Equal(true, obj.Remove(1))
	ast.Equal(false, obj.Insert(2))
	ast.Equal(2, obj.GetRandom())
}
