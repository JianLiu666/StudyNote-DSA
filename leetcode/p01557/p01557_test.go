package p01557

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	n     int
	edges [][]int
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
			input{n: 6, edges: [][]int{{0, 1}, {0, 2}, {2, 5}, {3, 4}, {4, 2}}},
			output{ans: []int{0, 3}},
		},
		{
			input{n: 5, edges: [][]int{{0, 1}, {2, 1}, {3, 1}, {1, 4}, {2, 4}}},
			output{ans: []int{0, 2, 3}},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, findSmallestSetOfVertices(data.i.n, data.i.edges), idx+1)
	}
}
