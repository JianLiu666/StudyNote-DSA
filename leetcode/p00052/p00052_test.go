package p00052

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	n int
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
			input{n: 4},
			output{ans: 2},
		},
		{
			input{n: 1},
			output{ans: 1},
		},
		{
			input{n: 9},
			output{ans: 352},
		},
		{
			input{n: 8},
			output{ans: 92},
		},
		{
			input{n: 5},
			output{ans: 10},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, totalNQueens(data.i.n), idx+1)
	}
}
