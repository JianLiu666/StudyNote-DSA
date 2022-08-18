package p01338

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	arr []int
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
			input{arr: []int{3, 3, 3, 3, 5, 5, 5, 2, 2, 7}},
			output{ans: 2},
		},
		{
			input{arr: []int{7, 7, 7, 7, 7, 7}},
			output{ans: 1},
		},
		{
			input{arr: []int{1, 9}},
			output{ans: 1},
		},
	}

	for idx, data := range tds {
		arr := make([]int, len(data.i.arr))
		copy(arr, data.i.arr)
		ast.Equal(data.o.ans, minSetSize_hash(arr), fmt.Sprintf("hash: %v", idx+1))

		arr = make([]int, len(data.i.arr))
		copy(arr, data.i.arr)
		ast.Equal(data.o.ans, minSetSize_sorting(arr), fmt.Sprintf("Sorting: %v", idx+1))
	}
}
