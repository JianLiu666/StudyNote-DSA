package p00347

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	nums []int
	k    int
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
			input{nums: []int{1, 1, 1, 2, 2, 3}, k: 2},
			output{ans: []int{1, 2}},
		},
		{
			input{nums: []int{1}, k: 1},
			output{ans: []int{1}},
		},
		{
			input{nums: []int{-1, -1}, k: 1},
			output{ans: []int{-1}},
		},
		{
			input{nums: []int{1, 2}, k: 2},
			output{ans: []int{1, 2}},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, topKFrequent(data.i.nums, data.i.k), idx+1)
	}
}
