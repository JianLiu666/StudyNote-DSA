package p00209

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	target int
	nums   []int
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
			input{target: 7, nums: []int{2, 3, 1, 2, 4, 3}},
			output{ans: 2},
		},
		{
			input{target: 4, nums: []int{1, 4, 4}},
			output{ans: 1},
		},
		{
			input{target: 11, nums: []int{1, 1, 1, 1, 1, 1, 1, 1}},
			output{ans: 0},
		},
		{
			input{target: 15, nums: []int{1, 2, 3, 4, 5}},
			output{ans: 5},
		},
	}

	for _, data := range tds {
		ast.Equal(data.o.ans, minSubArrayLen(data.i.target, data.i.nums))
	}
}
