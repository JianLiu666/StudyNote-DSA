package p00022

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	n int
}

type output struct {
	ans []string
}

type testdata struct {
	i input
	o output
}

func TestQuestion(t *testing.T) {
	ast := assert.New(t)

	tds := []testdata{
		{
			input{n: 3},
			output{ans: []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
		},
		{
			input{n: 1},
			output{ans: []string{"()"}},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, generateParenthesis(data.i.n), idx+1)
	}
}
