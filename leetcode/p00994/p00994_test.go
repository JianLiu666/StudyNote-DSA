package p00994

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	grid [][]int
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
			input{grid: [][]int{
				{2, 0},
				{0, 1},
			}},
			output{ans: -1},
		},
		{
			input{grid: [][]int{
				{2, 1, 1},
				{1, 1, 0},
				{0, 1, 1},
			}},
			output{ans: 4},
		},
		{
			input{grid: [][]int{
				{2, 1, 1},
				{0, 1, 1},
				{1, 0, 1},
			}},
			output{ans: -1},
		},
		{
			input{grid: [][]int{
				{0, 2},
			}},
			output{ans: 0},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, orangesRotting(data.i.grid), idx+1)
	}
}
