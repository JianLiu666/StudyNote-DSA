package p00017

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	digits string
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
			input{digits: "23"},
			output{ans: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		},
		{
			input{digits: ""},
			output{ans: []string{}},
		},
		{
			input{digits: "2"},
			output{ans: []string{"a", "b", "c"}},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, letterCombinations(data.i.digits), idx+1)
	}
}
