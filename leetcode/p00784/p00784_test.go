package p00784

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
			input{s: "a1b2"},
			output{ans: []string{"A1B2", "A1b2", "a1B2", "a1b2"}},
		},
		{
			input{s: "3z4"},
			output{ans: []string{"3Z4", "3z4"}},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, letterCasePermutation(data.i.s), idx+1)
	}
}
