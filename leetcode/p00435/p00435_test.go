package p00435

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	intervals [][]int
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
			input{intervals: [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}},
			output{ans: 1},
		},
		{
			input{intervals: [][]int{{1, 2}, {1, 2}, {1, 2}}},
			output{ans: 2},
		},
		{
			input{intervals: [][]int{{1, 2}, {2, 3}}},
			output{ans: 0},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, eraseOverlapIntervals(data.i.intervals), idx+1)
	}
}
