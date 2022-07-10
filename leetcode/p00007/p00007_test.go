package p00007

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	x int
}

type output struct {
	ans int
}

type testdata struct {
	i input
	o output
}

func TestQuestion(t *testing.T) {
	ast := assert.New(t)

	tds := []testdata{
		{
			input{x: 123},
			output{ans: 321},
		},
		{
			input{x: -123},
			output{ans: -321},
		},
		{
			input{x: 1234567899},
			output{ans: 0},
		},
	}

	for _, data := range tds {
		ast.Equal(data.o.ans, reverse_string(data.i.x))
		ast.Equal(data.o.ans, reverse_recursion(data.i.x))
	}
}
