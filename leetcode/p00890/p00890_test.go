package p00890

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	words   []string
	pattern string
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
			input{words: []string{"abc", "deq", "mee", "aqq", "dkd", "ccc"}, pattern: "abb"},
			output{ans: []string{"mee", "aqq"}},
		},
		{
			input{words: []string{"a", "b", "c"}, pattern: "a"},
			output{ans: []string{"a", "b", "c"}},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, findAndReplacePattern(data.i.words, data.i.pattern), idx+1)

	}
}
