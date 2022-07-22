package p00219

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	nums []int
	k    int
}

type output struct {
	ans bool
}

type testdata struct {
	i input
	o output
}

func TestQuestion(t *testing.T) {
	ast := assert.New(t)

	tds := []testdata{
		{
			input{nums: []int{1, 2, 3, 1}, k: 3},
			output{ans: true},
		},
		{
			input{nums: []int{1, 0, 1, 1}, k: 1},
			output{ans: true},
		},
		{
			input{nums: []int{1, 2, 3, 1, 2, 3}, k: 2},
			output{ans: false},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, containsNearbyDuplicate(data.i.nums, data.i.k), idx+1)
	}
}
