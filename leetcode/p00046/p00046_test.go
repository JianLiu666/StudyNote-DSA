package p00046

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
			input{nums: []int{1, 2, 3}},
			output{ans: [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}},
		},
		{
			input{nums: []int{0, 1}},
			output{ans: [][]int{{0, 1}, {1, 0}}},
		},
		{
			input{nums: []int{1}},
			output{ans: [][]int{{1}}},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, permute(data.i.nums), idx+1)
	}
}
