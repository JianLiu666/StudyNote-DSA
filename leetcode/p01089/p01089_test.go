package p01089

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	arr []int
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
			input{arr: []int{1, 0, 2, 3, 0, 4, 5, 0}},
			output{ans: []int{1, 0, 0, 2, 3, 0, 0, 4}},
		},
		{
			input{arr: []int{1, 2, 3}},
			output{ans: []int{1, 2, 3}},
		},
		{
			input{arr: []int{1, 5, 2, 0, 6, 8, 0, 6, 0}},
			output{ans: []int{1, 5, 2, 0, 0, 6, 8, 0, 0}},
		},
		{
			input{arr: []int{1, 5, 2, 0, 6, 8, 1, 0, 0}},
			output{ans: []int{1, 5, 2, 0, 0, 6, 8, 1, 0}},
		},
	}

	for _, data := range tds {
		duplicateZeros(data.i.arr)
		ast.Equal(data.o.ans, data.i.arr)
	}
}
