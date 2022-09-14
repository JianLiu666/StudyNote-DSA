package p00912

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	nums []int
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
			input{nums: []int{5, 2, 3, 1}},
			output{ans: []int{1, 2, 3, 5}},
		},
		{
			input{nums: []int{5, 1, 1, 2, 0, 0}},
			output{ans: []int{0, 0, 1, 1, 2, 5}},
		},
	}

	for idx, data := range tds {
		copyright := make([]int, len(data.i.nums))
		copy(copyright, data.i.nums)
		ast.Equal(data.o.ans, sortArray_heapSort(copyright), fmt.Sprintf("Heap Sort: %v", idx+1))

		copyright = make([]int, len(data.i.nums))
		copy(copyright, data.i.nums)
		ast.Equal(data.o.ans, sortArray_mergeSort_topDown(copyright), fmt.Sprintf("Merge Sort (Top-Down): %v", idx+1))

		copyright = make([]int, len(data.i.nums))
		copy(copyright, data.i.nums)
		ast.Equal(data.o.ans, sortArray_mergeSort_bottomUp(copyright), fmt.Sprintf("Merge Sort (Bottom-Up): %v", idx+1))
	}
}
