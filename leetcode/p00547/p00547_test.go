package p00547

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	isConnected [][]int
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
			input{isConnected: [][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}},
			output{ans: 2},
		},
		{
			input{isConnected: [][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}},
			output{ans: 3},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, findCircleNum(data.i.isConnected), idx+1)
	}
}
