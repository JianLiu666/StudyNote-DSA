package p00459

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	s string
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
			input{s: "abab"},
			output{ans: true},
		},
		{
			input{s: "aba"},
			output{ans: false},
		},
		{
			input{s: "abcabcabcabc"},
			output{ans: true},
		},
		{
			input{s: "babbabbabbabbab"},
			output{ans: true},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, repeatedSubstringPattern(data.i.s), idx+1)
	}
}
