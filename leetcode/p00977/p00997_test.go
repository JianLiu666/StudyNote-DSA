package p00977

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
			input{nums: []int{-4, -1, 0, 3, 10}},
			output{ans: []int{0, 1, 9, 16, 100}},
		},
		{
			input{nums: []int{-7, -3, 2, 3, 11}},
			output{ans: []int{4, 9, 9, 49, 121}},
		},
	}

	for _, data := range tds {
		ast.Equal(data.o.ans, sortedSquares(data.i.nums))
	}
}
