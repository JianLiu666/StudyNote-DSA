package p01202

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	s     string
	pairs [][]int
}

type output struct {
	ans string
}

type testdata struct {
	i input
	o output
}

func TestQuestion(t *testing.T) {
	ast := assert.New(t)

	tds := []testdata{
		{
			input{s: "dcab", pairs: [][]int{{0, 3}, {1, 2}}},
			output{ans: "bacd"},
		},
		{
			input{s: "dcab", pairs: [][]int{{0, 3}, {1, 2}, {0, 2}}},
			output{ans: "abcd"},
		},
		{
			input{s: "cba", pairs: [][]int{{0, 1}, {1, 2}}},
			output{ans: "abc"},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, smallestStringWithSwaps(data.i.s, data.i.pairs), idx+1)
	}
}
