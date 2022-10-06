package p00981

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuestion(t *testing.T) {
	ast := assert.New(t)

	obj := Constructor()
	obj.Set("love", "high", 10)
	obj.Set("love", "low", 20)

	ast.Equal("", obj.Get("love", 5), 1)
	ast.Equal("high", obj.Get("love", 10), 2)
	ast.Equal("high", obj.Get("love", 15), 3)
	ast.Equal("low", obj.Get("love", 20), 4)
	ast.Equal("low", obj.Get("love", 25), 5)
}
