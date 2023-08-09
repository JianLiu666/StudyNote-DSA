package p00525

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
			input{nums: []int{0, 1}},
			output{ans: 2},
		},
		{
			input{nums: []int{0, 1, 0}},
			output{ans: 2},
		},
		{
			input{nums: []int{0, 1, 0, 0, 1, 1, 0, 1, 0, 1}},
			output{ans: 10},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, findMaxLength(data.i.nums), idx+1)
	}
}
