package p00438

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	s string
	p string
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
			input{s: "cbaebabacd", p: "abc"},
			output{ans: []int{0, 6}},
		},
		{
			input{s: "abab", p: "ab"},
			output{ans: []int{0, 1, 2}},
		},
		{
			input{s: "aacbebabacd", p: "bbb"},
			output{ans: []int{}},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, findAnagrams(data.i.s, data.i.p), idx+1)
	}
}
