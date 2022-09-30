package p00035

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	nums   []int
	target int
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
			input{nums: []int{1, 3, 5, 6}, target: 5},
			output{ans: 2},
		},
		{
			input{nums: []int{1, 3, 5, 6}, target: 2},
			output{ans: 1},
		},
		{
			input{nums: []int{1, 3, 5, 6}, target: 7},
			output{ans: 4},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, searchInsert(data.i.nums, data.i.target), idx+1)
	}
}
