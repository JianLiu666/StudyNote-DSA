package p00409

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	s string
}

type output struct {
	ans int
}

type testdata struct {
	i input
	o output
}

func TestQuestion(t *testing.T) {
	ast := assert.New(t)

	tds := []testdata{
		{
			input{s: "abccccdd"},
			output{ans: 7},
		},
		{
			input{s: "a"},
			output{ans: 1},
		},
		{
			input{s: "ccc"},
			output{ans: 3},
		},
		{
			input{s: "aaaccc"},
			output{ans: 5},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, longestPalindrome(data.i.s), idx+1)
	}
}
