package p00151

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	s string
}

type output struct {
	ans string
}

type testdata struct {
	i input
	o output
}

func TestQuestion(t *testing.T) {
	ast := assert.New(t)

	tds := []testdata{
		{
			input{s: "the sky is blue"},
			output{ans: "blue is sky the"},
		},
		{
			input{s: "  hello world  "},
			output{ans: "world hello"},
		},
		{
			input{s: "a good   example"},
			output{ans: "example good a"},
		},
	}

	for _, data := range tds {
		ast.Equal(data.o.ans, reverseWords(data.i.s))
	}
}
