package p00744

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	letters []byte
	target  byte
}

type output struct {
	ans byte
}

type testdata struct {
	i input
	o output
}

func TestQuestion(t *testing.T) {
	ast := assert.New(t)

	tds := []testdata{
		{
			input{letters: []byte{'c', 'f', 'j'}, target: 'a'},
			output{ans: 'c'},
		},
		{
			input{letters: []byte{'c', 'f', 'j'}, target: 'c'},
			output{ans: 'f'},
		},
		{
			input{letters: []byte{'c', 'f', 'j'}, target: 'd'},
			output{ans: 'f'},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, nextGreatestLetter(data.i.letters, data.i.target), idx+1)
	}
}
