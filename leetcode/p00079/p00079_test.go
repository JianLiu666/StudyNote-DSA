package p00079

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	board [][]byte
	word  string
}

type output struct {
	ans bool
}

type testdata struct {
	i input
	o output
}

func TestQuestion(t *testing.T) {
	ast := assert.New(t)

	tds := []testdata{
		{
			input{word: "ABCCED", board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			}},
			output{
				ans: true,
			},
		},
		{
			input{word: "SEE", board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			}},
			output{
				ans: true,
			},
		},
		{
			input{word: "ABCB", board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			}},
			output{
				ans: false,
			},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, exist(data.i.board, data.i.word), idx+1)
	}
}
