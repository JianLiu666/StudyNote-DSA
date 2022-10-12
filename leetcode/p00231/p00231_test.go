package p00231

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
			input{n: 16},
			output{ans: true},
		},
		{
			input{n: 3},
			output{ans: false},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, isPowerOfTwo(data.i.n), idx+1)
	}
}
