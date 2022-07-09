package p00326

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	n int
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
			input{n: 27},
			output{ans: true},
		},
		{
			input{n: 0},
			output{ans: false},
		},
		{
			input{n: 9},
			output{ans: true},
		},
	}

	for _, data := range tds {
		ast.Equal(data.o.ans, isPowerOfThree(data.i.n))
	}
}
