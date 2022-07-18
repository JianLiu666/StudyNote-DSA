package p00189

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
			input{nums: []int{1, 2, 3, 4, 5, 6, 7}, k: 3},
			output{ans: []int{5, 6, 7, 1, 2, 3, 4}},
		},
		{
			input{nums: []int{-1, -100, 3, 99}, k: 2},
			output{ans: []int{3, 99, -1, -100}},
		},
	}

	for _, data := range tds {
		rotate(data.i.nums, data.i.k)
		ast.Equal(data.o.ans, data.i.nums)
	}
}
