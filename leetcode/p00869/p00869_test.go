package p00869

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
			input{n: 1},
			output{ans: true},
		},
		{
			input{n: 10},
			output{ans: false},
		},
		{
			input{n: 0},
			output{ans: false},
		},
		{
			input{n: 536870912},
			output{ans: true},
		},
		{
			input{n: 563789021},
			output{ans: true},
		},
		{
			input{n: 8},
			output{ans: true},
		},
		{
			input{n: 61},
			output{ans: true},
		},
		{
			input{n: 61},
			output{ans: true},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, reorderedPowerOf2(data.i.n), idx+1)
	}
}
