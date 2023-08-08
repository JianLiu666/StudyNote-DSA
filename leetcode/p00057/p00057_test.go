package p00057

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	intervals   [][]int
	newInterval []int
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
			input{
				intervals:   [][]int{{1, 3}, {6, 9}},
				newInterval: []int{2, 5},
			},
			output{
				ans: [][]int{{1, 5}, {6, 9}},
			},
		},
		{
			input{
				intervals:   [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
				newInterval: []int{4, 8},
			},
			output{
				ans: [][]int{{1, 2}, {3, 10}, {12, 16}},
			},
		},
		{
			input{
				intervals:   [][]int{},
				newInterval: []int{5, 7},
			},
			output{
				ans: [][]int{{5, 7}},
			},
		},
		{
			input{
				intervals:   [][]int{{1, 5}},
				newInterval: []int{6, 8},
			},
			output{
				ans: [][]int{{1, 5}, {6, 8}},
			},
		},
		{
			input{
				intervals:   [][]int{{1, 5}},
				newInterval: []int{0, 3},
			},
			output{
				ans: [][]int{{0, 5}},
			},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, insert(data.i.intervals, data.i.newInterval), idx+1)
	}
}
