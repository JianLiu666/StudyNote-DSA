package p00985

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	nums    []int
	queries [][]int
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
			input{
				nums:    []int{1, 2, 3, 4},
				queries: [][]int{{1, 0}, {-3, 1}, {-4, 0}, {2, 3}}},
			output{
				ans: []int{8, 6, 2, 4},
			},
		},
		{
			input{
				nums:    []int{1},
				queries: [][]int{{4, 0}}},
			output{
				ans: []int{0},
			},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, sumEvenAfterQueries(data.i.nums, data.i.queries), idx+1)
	}
}
