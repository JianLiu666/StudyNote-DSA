package p00394

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
			input{s: "abc"},
			output{ans: "abc"},
		},
		{
			input{s: "3[a]2[bc]"},
			output{ans: "aaabcbc"},
		},
		{
			input{s: "3[a]10[bc]"},
			output{ans: "aaabcbcbcbcbcbcbcbcbcbc"},
		},
		{
			input{s: "3[a2[c]]"},
			output{ans: "accaccacc"},
		},
		{
			input{s: "2[abc]3[cd]ef"},
			output{ans: "abcabccdcdcdef"},
		},
		{
			input{s: "3[z]2[2[y]pq4[2[jk]e1[f]]]ef"},
			output{ans: "zzzyypqjkjkefjkjkefjkjkefjkjkefyypqjkjkefjkjkefjkjkefjkjkefef"},
		},
		{
			input{s: "abc3[cd]xyz"},
			output{ans: "abccdcdcdxyz"},
		},
		{
			input{s: "sd2[f2[e]g]i"},
			output{ans: "sdfeegfeegi"},
		},
		{
			input{s: "1[a1[a1[a1[a1[a1[a1[a1[a1[a1[a1[a1[a1[a1[a]]]]]]]]]]]]]]"},
			output{ans: "aaaaaaaaaaaaaa"},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, decodeString(data.i.s), idx+1)
	}
}
