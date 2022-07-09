package p00344

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	s []byte
}

type output struct {
	ans []byte
}

type testdata struct {
	i input
	o output
}

func TestQuestion(t *testing.T) {
	ast := assert.New(t)

	tds := []testdata{
		{
			input{s: []byte{'h', 'e', 'l', 'l', 'o'}},
			output{ans: []byte{'o', 'l', 'l', 'e', 'h'}},
		},
		{
			input{s: []byte{'H', 'a', 'n', 'n', 'a', 'h'}},
			output{ans: []byte{'h', 'a', 'n', 'n', 'a', 'H'}},
		},
	}

	for _, data := range tds {
		reverseString(data.i.s)
		ast.Equal(data.o.ans, data.i.s)
	}
}
