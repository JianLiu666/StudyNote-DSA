package p00053

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
			input{nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}},
			output{ans: 6},
		},
		{
			input{nums: []int{1}},
			output{ans: 1},
		},
		{
			input{nums: []int{5, 4, -1, 7, 8}},
			output{ans: 23},
		},
	}

	for _, data := range tds {
		ast.Equal(data.o.ans, maxSubArray(data.i.nums))
	}
}
