package p00187

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	s string
}

type output struct {
	ans []string
}

type testdata struct {
	i input
	o output
}

func TestQuestion(t *testing.T) {
	ast := assert.New(t)

	tds := []testdata{
		{
			input{s: "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"},
			output{ans: []string{"AAAAACCCCC", "CCCCCAAAAA"}},
		},
		{
			input{s: "AAAAAAAAAAAAA"},
			output{ans: []string{"AAAAAAAAAA"}},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, findRepeatedDnaSequences(data.i.s), idx+1)
	}
}
