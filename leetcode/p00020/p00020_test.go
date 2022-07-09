package p00020

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	s string
}

type output struct {
	ans bool
}

type testdata struct {
	i input
	o output
}

func TestQuestion(t *testing.T) {
	ast := assert.New(t)

	tds := []testdata{
		{
			input{s: "()"},
			output{ans: true},
		},
		{
			input{s: "()[]{}"},
			output{ans: true},
		},
		{
			input{s: "(]"},
			output{ans: false},
		},
		{
			input{s: "([)]"},
			output{ans: false},
		},
	}

	for _, data := range tds {
		ast.Equal(data.o.ans, isValid(data.i.s))
	}
}
