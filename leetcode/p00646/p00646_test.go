package p00646

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	pairs [][]int
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
			input{
				pairs: [][]int{{1, 2}, {2, 3}, {3, 4}},
			},
			output{
				ans: 2,
			},
		},
		{
			input{
				pairs: [][]int{{1, 2}, {7, 8}, {4, 5}},
			},
			output{
				ans: 3,
			},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, findLongestChain(data.i.pairs), idx+1)
	}
}
