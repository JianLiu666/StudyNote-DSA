package p00205

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	s string
	t string
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
			input{s: "egg", t: "add"},
			output{ans: true},
		},
		{
			input{s: "foo", t: "bar"},
			output{ans: false},
		},
		{
			input{s: "paper", t: "title"},
			output{ans: true},
		},
		{
			input{s: "abcdefghijklmnopqrstuvwxyzva", t: "abcdefghijklmnopqrstuvwxyzck"},
			output{ans: false},
		},
		{
			input{s: "badc", t: "baba"},
			output{ans: false},
		},
	}

	for _, data := range tds {
		ast.Equal(data.o.ans, isIsomorphic(data.i.s, data.i.t))
	}
}
