package p00705

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuestion(t *testing.T) {
	ast := assert.New(t)

	obj := Constructor()
	obj.Add(1)
	obj.Remove(2)
	ast.Equal(true, obj.Contains(1))
	ast.Equal(false, obj.Contains(3))
	obj.Add(2)
	ast.Equal(true, obj.Contains(2))
	obj.Remove(2)
	ast.Equal(false, obj.Contains(2))
}
