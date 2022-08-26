package p00056

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	intervals [][]int
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
			input{intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}},
			output{ans: [][]int{{1, 6}, {8, 10}, {15, 18}}},
		},
		{
			input{intervals: [][]int{{1, 4}, {4, 5}}},
			output{ans: [][]int{{1, 5}}},
		},
		{
			input{intervals: [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}, {1, 10}}},
			output{ans: [][]int{{1, 10}}},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, merge(data.i.intervals), idx+1)
	}
}
