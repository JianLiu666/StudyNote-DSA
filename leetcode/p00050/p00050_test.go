package p00050

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	x float64
	n int
}

type output struct {
	ans float64
}

type testdata struct {
	i input
	o output
}

func TestQuestion(t *testing.T) {
	ast := assert.New(t)

	tds := []testdata{
		{
			input{x: 2.00000, n: 4},
			output{ans: 16.00000},
		},
		{
			input{x: 2.00000, n: 10},
			output{ans: 1024.00000},
		},
		{
			input{x: 2.10000, n: 3},
			output{ans: 9.26100},
		},
		{
			input{x: 2.00000, n: -2},
			output{ans: 0.25000},
		},
		{
			input{x: 34.00515, n: -3},
			output{ans: 0.00003},
		},
		{
			input{x: 8.84372, n: -5},
			output{ans: 0.00002},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, myPow(data.i.x, data.i.n), idx+1)
	}
}
