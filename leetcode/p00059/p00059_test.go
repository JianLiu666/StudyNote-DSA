package p00059

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	n int
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
			input{n: 3},
			output{ans: [][]int{
				{1, 2, 3},
				{8, 9, 4},
				{7, 6, 5},
			}},
		},
		{
			input{n: 1},
			output{ans: [][]int{
				{1},
			}},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, generateMatrix(data.i.n), idx+1)
	}
}
