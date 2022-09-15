package p00077

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	n int
	k int
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
			input{n: 4, k: 2},
			output{ans: [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}}},
		},
		{
			input{n: 1, k: 1},
			output{ans: [][]int{{1}}},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, combine(data.i.n, data.i.k), idx+1)
	}
}
