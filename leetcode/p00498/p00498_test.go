package p00498

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	mat [][]int
}

type output struct {
	ans []int
}

type testdata struct {
	i input
	o output
}

func TestQuestion(t *testing.T) {
	ast := assert.New(t)

	tds := []testdata{
		{
			input{mat: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}},
			output{ans: []int{1, 2, 4, 7, 5, 3, 6, 8, 9}},
		},
		{
			input{mat: [][]int{{1, 2}, {3, 4}}},
			output{ans: []int{1, 2, 3, 4}},
		},
		{
			input{mat: [][]int{{2, 3}}},
			output{ans: []int{2, 3}},
		},
	}

	for _, data := range tds {
		ast.Equal(data.o.ans, findDiagonalOrder(data.i.mat))
	}
}
