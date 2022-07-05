package p00033

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	nums   []int
	target int
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
			input{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 0},
			output{ans: 4},
		},
		{
			input{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 3},
			output{ans: -1},
		},
		{
			input{nums: []int{1}, target: 0},
			output{ans: -1},
		},
		{
			input{nums: []int{3, 1}, target: 1},
			output{ans: 1},
		},
	}

	for _, data := range tds {
		ast.Equal(data.o.ans, search(data.i.nums, data.i.target))
	}
}
