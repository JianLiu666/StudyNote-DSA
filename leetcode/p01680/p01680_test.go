package p01680

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
			input{n: 1},
			output{ans: 1},
		},
		{
			input{n: 3},
			output{ans: 27},
		},
		{
			input{n: 12},
			output{ans: 505379714},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, concatenatedBinary(data.i.n), idx+1)
	}
}
