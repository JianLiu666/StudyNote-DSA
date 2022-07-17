package p00167

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	numbers []int
	target  int
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
			input{numbers: []int{2, 7, 11, 15}, target: 9},
			output{ans: []int{1, 2}},
		},
		{
			input{numbers: []int{2, 3, 4}, target: 6},
			output{ans: []int{1, 3}},
		},
		{
			input{numbers: []int{-1, 0}, target: -1},
			output{ans: []int{1, 2}},
		},
	}

	for _, data := range tds {
		ast.Equal(data.o.ans, twoSum(data.i.numbers, data.i.target))
	}
}
