package p00215

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	nums []int
	k    int
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
			input{nums: []int{3, 2, 1, 5, 6, 4}, k: 2},
			output{ans: 5},
		},
		{
			input{nums: []int{3, 2, 3, 1, 2, 4, 5, 5, 6}, k: 4},
			output{ans: 4},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, findKthLargest_heapsort(data.i.nums, data.i.k), fmt.Sprintf("Heap Sort: %v", idx+1))
	}
}
