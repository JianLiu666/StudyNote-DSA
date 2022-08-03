package p00287

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	nums []int
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
			input{nums: []int{1, 3, 4, 2, 2}},
			output{ans: 2},
		},
		{
			input{nums: []int{3, 1, 3, 4, 2}},
			output{ans: 3},
		},
		{
			input{nums: []int{1, 3, 4, 2, 2}},
			output{ans: 2},
		},
		{
			input{nums: []int{3, 1, 3, 4, 2}},
			output{ans: 3},
		},
		{
			input{nums: []int{1, 1, 2, 3, 4}},
			output{ans: 1},
		},
		{
			input{nums: []int{1, 2, 2, 3, 4}},
			output{ans: 2},
		},
		{
			input{nums: []int{1, 2, 4, 4, 4}},
			output{ans: 4},
		},
		{
			input{nums: []int{4, 1, 2, 4, 3}},
			output{ans: 4},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, findDuplicate_binarysearch(data.i.nums), fmt.Sprintf("Binary Search: %v", idx+1))
		ast.Equal(data.o.ans, findDuplicate_cycledetection(data.i.nums), fmt.Sprintf("Cycle Detection: %v", idx+1))
	}
}
