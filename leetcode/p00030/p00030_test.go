package p00030

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	s     string
	words []string
}

type output struct {
	ans []int
}

type testdata struct {
	i input
	o output
}

func TestQuestion(t *testing.T) {
	ast := assert.New(t)

	tds := []testdata{
		{
			input{s: "barfoothefoobarman", words: []string{"foo", "bar"}},
			output{ans: []int{0, 9}},
		},
		{
			input{s: "wordgoodgoodgoodbestword", words: []string{"word", "good", "best", "word"}},
			output{ans: []int{}},
		},
		{
			input{s: "barfoofoobarthefoobarman", words: []string{"bar", "foo", "the"}},
			output{ans: []int{6, 9, 12}},
		},
		{
			input{s: "aaaaaaaaaaaaaa", words: []string{"aa", "aa"}},
			output{ans: []int{0, 2, 4, 6, 8, 10, 1, 3, 5, 7, 9}},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, findSubstring(data.i.s, data.i.words), idx+1)
	}
}
