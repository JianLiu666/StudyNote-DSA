package p02309

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	s string
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
			input{s: "lEeTcOdE"},
			output{ans: "E"},
		},
		{
			input{s: "arRAzFif"},
			output{ans: "R"},
		},
		{
			input{s: "AbCdEfGhIjK"},
			output{ans: ""},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, greatestLetter(data.i.s), idx+1)
	}
}
