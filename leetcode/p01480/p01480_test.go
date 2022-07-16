package p01480

import (
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
			input{nums: []int{1, 2, 3, 4}},
			output{ans: []int{1, 3, 6, 10}},
		},
		{
			input{nums: []int{1, 1, 1, 1, 1}},
			output{ans: []int{1, 2, 3, 4, 5}},
		},
		{
			input{nums: []int{3, 1, 2, 10, 1}},
			output{ans: []int{3, 4, 6, 16, 17}},
		},
	}

	for _, data := range tds {
		ast.Equal(data.o.ans, runningSum(data.i.nums))
	}
}
