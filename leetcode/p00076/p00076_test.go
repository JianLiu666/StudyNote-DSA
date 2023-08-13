package p00076

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	s string
	t string
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
			input{s: "ADOBECODEBANC", t: "ABC"},
			output{ans: "BANC"},
		},
		{
			input{s: "a", t: "a"},
			output{ans: "a"},
		},
		{
			input{s: "a", t: "aa"},
			output{ans: ""},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, data.i.s, data.i.t, idx+1)
	}
}
