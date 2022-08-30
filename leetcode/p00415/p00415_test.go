package p00415

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	num1 string
	num2 string
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
			input{num1: "11", num2: "123"},
			output{ans: "134"},
		},
		{
			input{num1: "9", num2: "99"},
			output{ans: "108"},
		},
		{
			input{num1: "456", num2: "77"},
			output{ans: "533"},
		},
		{
			input{num1: "1", num2: "9"},
			output{ans: "10"},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, addStrings(data.i.num1, data.i.num2), idx+1)
	}
}
