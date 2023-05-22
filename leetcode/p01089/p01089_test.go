package p01089

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	arr []int
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
			input{arr: []int{1, 0, 2, 3, 0, 4, 5, 0}},
			output{ans: []int{1, 0, 0, 2, 3, 0, 0, 4}},
		},
		{
			input{arr: []int{1, 2, 3}},
			output{ans: []int{1, 2, 3}},
		},
		{
			input{arr: []int{1, 5, 2, 0, 6, 8, 0, 6, 0}},
			output{ans: []int{1, 5, 2, 0, 0, 6, 8, 0, 0}},
		},
		{
			input{arr: []int{1, 5, 2, 0, 6, 8, 1, 0, 0}},
			output{ans: []int{1, 5, 2, 0, 0, 6, 8, 1, 0}},
		},
		{
			input{arr: []int{8, 4, 5, 0, 0, 0, 7}},
			output{ans: []int{8, 4, 5, 0, 0, 0, 0}},
		},
		{
			input{arr: []int{0, 0, 0, 0, 0, 0, 0}},
			output{ans: []int{0, 0, 0, 0, 0, 0, 0}},
		},
		{
			input{arr: []int{8, 4, 5, 0, 0, 0, 0, 7}},
			output{ans: []int{8, 4, 5, 0, 0, 0, 0, 0}},
		},
	}

	for idx, data := range tds {
		copyright := make([]int, len(data.i.arr))
		copy(copyright, data.i.arr)
		duplicateZeros_sc_o1(copyright)
		ast.Equal(data.o.ans, copyright, fmt.Sprintf("Space Complexity O(1): %v", idx+1))

		copyright = make([]int, len(data.i.arr))
		copy(copyright, data.i.arr)
		duplicateZeros_sc_on(copyright)
		ast.Equal(data.o.ans, copyright, fmt.Sprintf("Space Complexity O(n): %v", idx+1))
	}
}
