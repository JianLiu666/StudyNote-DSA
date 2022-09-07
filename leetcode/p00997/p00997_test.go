package p00997

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	n     int
	trust [][]int
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
			input{n: 2, trust: [][]int{{1, 2}}},
			output{ans: 2},
		},
		{
			input{n: 3, trust: [][]int{{1, 3}, {2, 3}}},
			output{ans: 3},
		},
		{
			input{n: 3, trust: [][]int{{1, 3}, {2, 3}, {3, 1}}},
			output{ans: -1},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, findJudge_hash(data.i.n, data.i.trust), fmt.Sprintf("Hash Table: %v", idx+1))
		ast.Equal(data.o.ans, findJudge_graph(data.i.n, data.i.trust), fmt.Sprintf("Graph: %v", idx+1))
	}
}
