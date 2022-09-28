package p00128

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
			input{nums: []int{100, 4, 200, 1, 3, 2}},
			output{ans: 4},
		},
		{
			input{nums: []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}},
			output{ans: 9},
		},
		{
			input{nums: []int{1, 0, 1, 2}},
			output{ans: 3},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, longestConsecutive(data.i.nums), idx+1)
	}
}
