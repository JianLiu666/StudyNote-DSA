package p01832

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	sentence string
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
			input{sentence: "thequickbrownfoxjumpsoverthelazydog"},
			output{ans: true},
		},
		{
			input{sentence: "leetcode"},
			output{ans: false},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, checkIfPangram(data.i.sentence), idx+1)
	}
}
