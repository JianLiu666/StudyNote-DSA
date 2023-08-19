package p00139

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	s        string
	wordDict []string
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
			input{
				s:        "leetcode",
				wordDict: []string{"leet", "code"},
			},
			output{
				ans: true,
			},
		},
		{
			input{
				s:        "applepenapple",
				wordDict: []string{"apple", "pen"},
			},
			output{
				ans: true,
			},
		},
		{
			input{
				s:        "catsandog",
				wordDict: []string{"cats", "dog", "sand", "and", "cat"},
			},
			output{
				ans: false,
			},
		},
		{
			input{
				s:        "aaaaaaa",
				wordDict: []string{"aaaa", "aaa"},
			},
			output{
				ans: true,
			},
		},
		{
			input{
				s:        "ab",
				wordDict: []string{"a", "b"},
			},
			output{
				ans: true,
			},
		},
		{
			input{
				s:        "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab",
				wordDict: []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"},
			},
			output{
				ans: false,
			},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, wordBreak(data.i.s, data.i.wordDict), idx+1)
	}
}
