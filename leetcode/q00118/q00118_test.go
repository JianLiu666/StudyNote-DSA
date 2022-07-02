package q00118

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	numRows int
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
			input{numRows: 5},
			output{ans: [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}}},
		},
		{
			input{numRows: 1},
			output{ans: [][]int{{1}}},
		},
	}

	for _, data := range tds {
		ast.Equal(data.o.ans, generate(data.i.numRows))
	}
}
