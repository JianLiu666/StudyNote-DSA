package p00153

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
			input{nums: []int{3, 4, 5, 1, 2}},
			output{ans: 1},
		},
		{
			input{nums: []int{4, 5, 6, 7, 0, 1, 2}},
			output{ans: 0},
		},
		{
			input{nums: []int{11, 13, 15, 17}},
			output{ans: 11},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, findMin(data.i.nums), idx+1)
	}
}
