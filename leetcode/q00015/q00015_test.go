package q00015

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	nums []int
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
			input{nums: []int{-4, -1, -1, 0, 1, 2}},
			output{ans: [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		},
		{
			input{nums: []int{-4, -3, -2, 6, 7}},
			output{ans: [][]int{{-4, -3, 7}, {-4, -2, 6}}},
		},
		{
			input{nums: []int{}},
			output{ans: [][]int{}},
		},
		{
			input{nums: []int{0}},
			output{ans: [][]int{}},
		},
		{
			input{nums: []int{0, 0, 0, 0}},
			output{ans: [][]int{{0, 0, 0}}},
		},
	}

	for _, data := range tds {
		ast.Equal(data.o.ans, threeSum(data.i.nums))
	}
}
