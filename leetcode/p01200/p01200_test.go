package p01200

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	arr []int
}

type output struct {
	ans [][]int
}

type testdata struct {
	i input
	o output
}

func TestQuestion(t *testing.T) {
	ast := assert.New(t)

	tds := []testdata{
		{
			input{arr: []int{4, 2, 1, 3}},
			output{ans: [][]int{{1, 2}, {2, 3}, {3, 4}}},
		},
		{
			input{arr: []int{1, 3, 6, 10, 15}},
			output{ans: [][]int{{1, 3}}},
		},
		{
			input{arr: []int{3, 8, -10, 23, 19, -4, -14, 27}},
			output{ans: [][]int{{-14, -10}, {19, 23}, {23, 27}}},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, minimumAbsDifference(data.i.arr), idx+1)
	}
}
