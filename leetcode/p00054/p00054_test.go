package p00054

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	matrix [][]int
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
			input{matrix: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}},
			output{ans: []int{1, 2, 3, 6, 9, 8, 7, 4, 5}},
		},
		{
			input{matrix: [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}},
			output{ans: []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}},
		},
	}

	for _, data := range tds {
		ast.Equal(data.o.ans, spiralOrder(data.i.matrix))
	}
}
