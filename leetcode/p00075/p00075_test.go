package p00075

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
			input{nums: []int{2, 0, 2, 1, 1, 0}},
			output{ans: []int{0, 0, 1, 1, 2, 2}},
		},
		{
			input{nums: []int{2, 0, 1}},
			output{ans: []int{0, 1, 2}},
		},
	}

	for idx, data := range tds {
		sortColors(data.i.nums)
		ast.Equal(data.o.ans, data.i.nums, idx+1)
	}
}
