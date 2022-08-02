package p00378

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	matrix [][]int
	k      int
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
			input{matrix: [][]int{{1, 5, 9}, {10, 11, 13}, {12, 13, 15}}, k: 8},
			output{ans: 13},
		},
		{
			input{matrix: [][]int{{-5}}, k: 1},
			output{ans: -5},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, kthSmallest(data.i.matrix, data.i.k), idx+1)
	}
}
