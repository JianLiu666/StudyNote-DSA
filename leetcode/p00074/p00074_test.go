package p00074

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	matrix [][]int
	target int
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
			input{
				matrix: [][]int{
					{1, 3, 5, 7},
					{10, 11, 16, 20},
					{23, 30, 34, 60}},
				target: 3},
			output{ans: true},
		},
		{
			input{
				matrix: [][]int{
					{1, 3, 5, 7},
					{10, 11, 16, 20},
					{23, 30, 34, 60}},
				target: 13},
			output{ans: false},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, searchMatrix(data.i.matrix, data.i.target), idx+1)
	}
}
