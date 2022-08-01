package p00704

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
			input{nums: []int{-1, 0, 3, 5, 9, 12}, target: 9},
			output{ans: 4},
		},
		{
			input{nums: []int{-1, 0, 3, 5, 9, 12}, target: 2},
			output{ans: -1},
		},
	}

	for _, data := range tds {
		ast.Equal(data.o.ans, search(data.i.nums, data.i.target))
	}
}
