package p00239

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	nums []int
	k    int
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
			input{nums: []int{1, 3, -1, -3, 5, 3, 6, 7}, k: 3},
			output{ans: []int{3, 3, 5, 5, 6, 7}},
		},
		{
			input{nums: []int{1}, k: 1},
			output{ans: []int{1}},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, maxSlidingWindow(data.i.nums, data.i.k), idx+1)
	}
}
