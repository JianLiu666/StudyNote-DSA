package q00268

import (
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
			input{nums: []int{3, 0, 1}},
			output{ans: 2},
		},
		{
			input{nums: []int{0, 1}},
			output{ans: 2},
		},
		{
			input{nums: []int{9, 6, 4, 2, 3, 5, 7, 0, 1}},
			output{ans: 8},
		},
	}

	for _, data := range tds {
		ast.Equal(data.o.ans, missingNumber_xor(data.i.nums))
		ast.Equal(data.o.ans, missingNumber_math(data.i.nums))
		ast.Equal(data.o.ans, missingNumber_map(data.i.nums))
		ast.Equal(data.o.ans, missingNumber_sort(data.i.nums))
	}
}

func TestMergeSort(t *testing.T) {
	ast := assert.New(t)

	ast.Equal([]int{0, 1, 3}, merge_sort([]int{3, 0, 1}))
	ast.Equal([]int{0, 1}, merge_sort([]int{0, 1}))
	ast.Equal([]int{0, 1, 2, 3, 4, 5, 6, 7, 9}, merge_sort([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
}
