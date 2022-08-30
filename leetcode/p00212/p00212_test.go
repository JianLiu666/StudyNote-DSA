package p00212

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	board [][]byte
	words []string
}

type output struct {
	ans []string
}

type testdata struct {
	i input
	o output
}

func TestQuestion(t *testing.T) {
	ast := assert.New(t)

	tds := []testdata{
		{
			input{
				board: [][]byte{
					{'o', 'a', 'a', 'n'},
					{'e', 't', 'a', 'e'},
					{'i', 'h', 'k', 'r'},
					{'i', 'f', 'l', 'v'}},
				words: []string{"oath", "pea", "eat", "rain"},
			},
			output{
				ans: []string{"eat", "oath"},
			},
		},
		{
			input{
				board: [][]byte{
					{'a', 'b'},
					{'c', 'd'}},
				words: []string{"abcb"},
			},
			output{
				ans: []string{},
			},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, findWords(data.i.board, data.i.words), idx+1)
	}
}
