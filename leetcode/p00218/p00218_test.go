package p00218

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	buildings [][]int
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
			input{buildings: [][]int{
				{2, 9, 10},
				{3, 7, 15},
				{5, 12, 12},
				{15, 20, 10},
				{19, 24, 8},
			}},
			output{ans: [][]int{
				{2, 10},
				{3, 15},
				{7, 12},
				{12, 0},
				{15, 10},
				{20, 8},
				{24, 0},
			}},
		},
		{
			input{buildings: [][]int{
				{0, 2, 3},
				{2, 5, 3},
			}},
			output{ans: [][]int{
				{0, 3},
				{5, 0},
			}},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, getSkyline(data.i.buildings), idx+1)
	}
}
