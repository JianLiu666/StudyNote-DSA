package p00200

import (
	"fmt"
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
		copyright := make([][]byte, len(data.i.grid))
		for i := 0; i < len(data.i.grid); i++ {
			arr := make([]byte, len(data.i.grid[i]))
			copy(arr, data.i.grid[i])
			copyright[i] = arr
		}
		ast.Equal(data.o.ans, numIslands_bfs(copyright), fmt.Sprintf("bfs case %d", idx+1))

		copyright = make([][]byte, len(data.i.grid))
		for i := 0; i < len(data.i.grid); i++ {
			arr := make([]byte, len(data.i.grid[i]))
			copy(arr, data.i.grid[i])
			copyright[i] = arr
		}
		ast.Equal(data.o.ans, numIslands_dfs(copyright), fmt.Sprintf("dfs case %d", idx+1))
	}
}
