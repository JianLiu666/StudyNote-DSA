package p00724

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
			input{nums: []int{1, 7, 3, 6, 5, 6}},
			output{ans: 3},
		},
		{
			input{nums: []int{1, 2, 3}},
			output{ans: -1},
		},
		{
			input{nums: []int{2, 1, -1}},
			output{ans: 0},
		},
	}

	for _, data := range tds {
		ast.Equal(data.o.ans, pivotIndex(data.i.nums))
	}
}
