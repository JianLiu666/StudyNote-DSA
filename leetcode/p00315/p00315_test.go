package p00315

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
			input{nums: []int{5, 2, 6, 1}},
			output{ans: []int{2, 1, 1, 0}},
		},
		{
			input{nums: []int{-1}},
			output{ans: []int{0}},
		},
		{
			input{nums: []int{-1, -1}},
			output{ans: []int{0, 0}},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, countSmaller(data.i.nums), idx+1)
	}
}
