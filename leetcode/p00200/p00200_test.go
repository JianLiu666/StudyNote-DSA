package p00200

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	grid [][]byte
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
			input{grid: [][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			}},
			output{ans: 1},
		},
		{
			input{grid: [][]byte{
				{'1', '1', '1', '0', '0'},
				{'1', '0', '1', '0', '0'},
				{'1', '1', '1', '0', '0'},
				{'0', '0', '0', '0', '0'},
			}},
			output{ans: 1},
		},
		{
			input{grid: [][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			}},
			output{ans: 3},
		},
		{
			input{grid: [][]byte{
				{'1', '1', '1'},
				{'0', '1', '0'},
				{'1', '1', '1'},
			}},
			output{ans: 1},
		},
		{
			input{grid: [][]byte{
				{'1', '1', '1', '1'},
				{'0', '1', '0', '0'},
				{'0', '1', '0', '1'},
				{'1', '1', '1', '1'},
			}},
			output{ans: 1},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, numIslands(data.i.grid), idx+1)
	}
}
