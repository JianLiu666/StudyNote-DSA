package p00424

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	s string
	k int
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
			input{s: "ABAB", k: 2},
			output{ans: 4},
		},
		{
			input{s: "AABABBA", k: 1},
			output{ans: 4},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, characterReplacement(data.i.s, data.i.k), idx+1)
	}
}
