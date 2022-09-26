package p02343

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	nums    []string
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
				nums:    []string{"102", "473", "251", "814"},
				queries: [][]int{{1, 1}, {2, 3}, {4, 2}, {1, 2}}},
			output{
				ans: []int{2, 2, 1, 0},
			},
		},
		{
			input{
				nums:    []string{"24", "37", "96", "04"},
				queries: [][]int{{2, 1}, {2, 2}}},
			output{
				ans: []int{3, 0},
			},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, smallestTrimmedNumbers(data.i.nums, data.i.queries), idx+1)
	}
}
