package q00001

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
			input{nums: []int{2, 7, 11, 15}, target: 9},
			output{ans: []int{0, 1}},
		},
		{
			input{nums: []int{3, 2, 4}, target: 6},
			output{ans: []int{1, 2}},
		},
		{
			input{nums: []int{3, 3}, target: 6},
			output{ans: []int{0, 1}},
		},
	}

	for _, data := range tds {
		ast.Equal(data.o.ans, twoSum(data.i.nums, data.i.target))
	}
}
