package p00013

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	s string
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
			input{s: "III"},
			output{ans: 3},
		},
		{
			input{s: "LVIII"},
			output{ans: 58},
		},
		{
			input{s: "MCMXCIV"},
			output{ans: 1994},
		},
	}

	for _, data := range tds {
		ast.Equal(data.o.ans, romanToInt(data.i.s))
	}
}
