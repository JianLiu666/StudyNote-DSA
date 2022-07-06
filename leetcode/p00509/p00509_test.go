package p00509

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	n int
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
			input{n: 2},
			output{ans: 1},
		},
		{
			input{n: 3},
			output{ans: 2},
		},
		{
			input{n: 4},
			output{ans: 3},
		},
	}

	for _, data := range tds {
		ast.Equal(data.o.ans, fib_recursion(data.i.n))
		ast.Equal(data.o.ans, fib_dp(data.i.n))
		ast.Equal(data.o.ans, fib_math(data.i.n))
	}
}
