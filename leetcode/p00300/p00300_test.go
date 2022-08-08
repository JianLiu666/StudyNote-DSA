package p00300

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
			input{nums: []int{10, 9, 2, 5, 3, 7, 101, 18}},
			output{ans: 4},
		},
		{
			input{nums: []int{0, 1, 0, 3, 2, 3}},
			output{ans: 4},
		},
		{
			input{nums: []int{7, 7, 7, 7, 7, 7, 7, 7}},
			output{ans: 1},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, lengthOfLIS_binarysearch(data.i.nums), fmt.Sprintf("Binary Search: %v", idx+1))
		ast.Equal(data.o.ans, lengthOfLIS_dp(data.i.nums), fmt.Sprintf("Dynamic Programming: %v", idx+1))
	}
}
