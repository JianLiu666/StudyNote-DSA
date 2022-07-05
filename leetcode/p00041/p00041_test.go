package p00041

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
			input{nums: []int{1, 2, 0}},
			output{ans: 3},
		},
		{
			input{nums: []int{1, 2, 3}},
			output{ans: 4},
		},
		{
			input{nums: []int{3, 4, -1, 1}},
			output{ans: 2},
		},
	}

	for _, data := range tds {
		ast.Equal(data.o.ans, firstMissingPositive_swap(data.i.nums))
		ast.Equal(data.o.ans, firstMissingPositive_hash(data.i.nums))
	}
}
