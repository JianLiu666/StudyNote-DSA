package p00566

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	mat [][]int
	r   int
	c   int
}

type output struct {
	ans [][]int
}

type testdata struct {
	i input
	o output
}

func TestQuestion(t *testing.T) {
	ast := assert.New(t)

	tds := []testdata{
		{
			input{mat: [][]int{{1, 2}, {3, 4}}, r: 2, c: 4},
			output{ans: [][]int{{1, 2}, {3, 4}}},
		},
		{
			input{mat: [][]int{{1, 2}, {3, 4}}, r: 1, c: 4},
			output{ans: [][]int{{1, 2, 3, 4}}},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, matrixReshape(data.i.mat, data.i.r, data.i.c), idx+1)
	}
}
