package p00719

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	nums []int
	k    int
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
			input{nums: []int{1, 3, 1}, k: 1},
			output{ans: 0},
		},
		{
			input{nums: []int{1, 1, 1}, k: 2},
			output{ans: 0},
		},
		{
			input{nums: []int{1, 6, 1}, k: 3},
			output{ans: 5},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, smallestDistancePair(data.i.nums, data.i.k), idx+1)
	}
}
