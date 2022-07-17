package p00067

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	a string
	b string
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
			input{a: "11", b: "1"},
			output{ans: "100"},
		},
		{
			input{a: "1010", b: "1011"},
			output{ans: "10101"},
		},
	}

	for _, data := range tds {
		ast.Equal(data.o.ans, addBinary_bruteforce(data.i.a, data.i.b))
		ast.Equal(data.o.ans, addBinary_builtin(data.i.a, data.i.b))
	}
}
