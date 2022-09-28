package p00078

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
			output{ans: [][]int{{}, {3}, {2}, {2, 3}, {1}, {1, 3}, {1, 2}, {1, 2, 3}}},
		},
		{
			input{nums: []int{0}},
			output{ans: [][]int{{}, {0}}},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, subsets(data.i.nums), idx+1)
	}
}
