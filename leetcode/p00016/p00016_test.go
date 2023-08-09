package p00016

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
			input{nums: []int{-1, 2, 1, -4}, target: 1},
			output{ans: 2},
		},
		{
			input{nums: []int{0, 0, 0}, target: 1},
			output{ans: 0},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, threeSumClosest(data.i.nums, data.i.target), idx+1)
	}
}
