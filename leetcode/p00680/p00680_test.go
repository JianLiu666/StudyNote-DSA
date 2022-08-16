package p00680

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
			input{s: "aba"},
			output{ans: true},
		},
		{
			input{s: "abca"},
			output{ans: true},
		},
		{
			input{s: "abc"},
			output{ans: false},
		},
		{
			input{s: "aguokepatgbnvfqmgmlcupuufxoohdfpgjdmysgvhmvffcnqxjjxqncffvmhvgsymdjgpfdhooxfuupuculmgmqfvnbgtapekouga"},
			output{ans: true},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, validPalindrome(data.i.s), idx+1)
	}
}
