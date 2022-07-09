package p00069

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	x int
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
			input{x: 4},
			output{ans: 2},
		},
		{
			input{x: 8},
			output{ans: 2},
		},
		{
			input{x: 808909662},
			output{ans: 28441},
		},
	}

	for _, data := range tds {
		ast.Equal(data.o.ans, mySqrt(data.i.x))
	}
}
