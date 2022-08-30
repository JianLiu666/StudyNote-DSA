package p00212

import (
	"sort"
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
					{'o', 'a', 'b', 'n'},
					{'o', 't', 'a', 'e'},
					{'a', 'h', 'k', 'r'},
					{'a', 'f', 'l', 'v'}},
				words: []string{"oa", "oaa"},
			},
			output{
				ans: []string{"oa", "oaa"},
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
		{
			input{
				board: [][]byte{
					{'a', 'a'}},
				words: []string{"aaa"},
			},
			output{
				ans: []string{},
			},
		},
		{
			input{
				board: [][]byte{
					{'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a'},
					{'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a'},
					{'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a'},
					{'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a'},
					{'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a'},
					{'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a'},
					{'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a'},
					{'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a'},
					{'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a'},
					{'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a'},
					{'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a'},
					{'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a'}},
				words: []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"},
			},
			output{
				ans: []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"},
			},
		},
	}

	for idx, data := range tds {
		res := findWords(data.i.board, data.i.words)
		sort.Strings(res)
		ast.Equal(data.o.ans, res, idx+1)
	}
}
