package p00202

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
			input{n: 19},
			output{ans: true},
		},
		{
			input{n: 2},
			output{ans: false},
		},
	}

	for _, data := range tds {
		ast.Equal(data.o.ans, isHappy(data.i.n))
		ast.Equal(data.o.ans, isHappy_twoPointers(data.i.n))
		ast.Equal(data.o.ans, isHappy_math(data.i.n))
	}
}
