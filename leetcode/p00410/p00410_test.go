package p00410

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	nums []int
	m    int
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
			input{nums: []int{7, 2, 5, 10, 8}, m: 2},
			output{ans: 18},
		},
		{
			input{nums: []int{1, 2, 3, 4, 5}, m: 2},
			output{ans: 9},
		},
		{
			input{nums: []int{1, 4, 4}, m: 3},
			output{ans: 4},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, splitArray(data.i.nums, data.i.m), idx+1)
	}
}
