package p00290

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	pattern string
	s       string
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
			input{pattern: "abba", s: "dog cat cat dog"},
			output{ans: true},
		},
		{
			input{pattern: "abba", s: "dog cat cat fish"},
			output{ans: false},
		},
		{
			input{pattern: "aaaa", s: "dog cat cat dog"},
			output{ans: false},
		},
		{
			input{pattern: "aaa", s: "aa aa aa aa"},
			output{ans: false},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, wordPattern(data.i.pattern, data.i.s), idx+1)
	}
}
