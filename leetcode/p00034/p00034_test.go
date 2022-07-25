package p00034

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	nums   []int
	target int
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
			input{nums: []int{5, 7, 7, 8, 8, 8, 8, 8, 8, 9, 9, 9}, target: 8},
			output{ans: []int{3, 8}},
		},
		{
			input{nums: []int{5, 7, 7, 8, 8, 10}, target: 8},
			output{ans: []int{3, 4}},
		},
		{
			input{nums: []int{5, 7, 7, 8, 8, 10}, target: 6},
			output{ans: []int{-1, -1}},
		},
		{
			input{nums: []int{}, target: 0},
			output{ans: []int{-1, -1}},
		},
		{
			input{nums: []int{2, 2}, target: 2},
			output{ans: []int{0, 1}},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, searchRange(data.i.nums, data.i.target), idx+1)
	}
}
