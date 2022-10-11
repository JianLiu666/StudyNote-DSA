package p00120

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	triangle [][]int
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
			input{triangle: [][]int{
				{2},
				{3, 4},
				{6, 5, 7},
				{4, 1, 8, 3},
			}},
			output{ans: 11},
		},
		{
			input{triangle: [][]int{
				{-10},
			}},
			output{ans: -10},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, minimumTotal(data.i.triangle), idx+1)
	}
}
