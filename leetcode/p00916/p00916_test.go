package p00916

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	words1 []string
	words2 []string
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
				words1: []string{"amazon", "apple", "facebook", "google", "leetcode"},
				words2: []string{"e", "o"}},
			output{
				ans: []string{"facebook", "google", "leetcode"},
			},
		},
		{
			input{
				words1: []string{"amazon", "apple", "facebook", "google", "leetcode"},
				words2: []string{"l", "e"}},
			output{
				ans: []string{"apple", "google", "leetcode"},
			},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, wordSubsets(data.i.words1, data.i.words2), idx+1)

	}
}
