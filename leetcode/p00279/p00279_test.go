package p00279

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
			input{n: 12},
			output{ans: 3},
		},
		{
			input{n: 13},
			output{ans: 2},
		},
		{
			input{n: 8328},
			output{ans: 3},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, numSquares(data.i.n), idx+1)
	}
}
