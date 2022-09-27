package p00838

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	dominoes string
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
			input{dominoes: "RR.L"},
			output{ans: "RR.L"},
		},
		{
			input{dominoes: ".L.R...LR..L.."},
			output{ans: "LL.RR.LLRRLL.."},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, pushDominoes(data.i.dominoes), idx+1)
	}
}
