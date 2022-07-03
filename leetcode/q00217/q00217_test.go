package q00217

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	nums []int
}

type output struct {
	ans bool
}

type testdata struct {
	i input
	o output
}

func TestQuestion(t *testing.T) {
	ast := assert.New(t)

	tds := []testdata{
		{
			input{nums: []int{1, 2, 3, 1}},
			output{ans: true},
		},
		{
			input{nums: []int{1, 2, 3, 4}},
			output{ans: false},
		},
		{
			input{nums: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}},
			output{ans: true},
		},
		{
			input{nums: []int{3, 1}},
			output{ans: false},
		},
	}

	for _, data := range tds {
		ast.Equal(data.o.ans, containsDuplicate_map(data.i.nums))
		ast.Equal(data.o.ans, containsDuplicate_sort(data.i.nums))
	}
}

func TestMergeSort(t *testing.T) {
	fmt.Println(merge_sort([]int{1, 2, 3, 1}))
	fmt.Println(merge_sort([]int{1, 2, 3, 4}))
}
