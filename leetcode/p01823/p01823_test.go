package p01823

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	n int
	k int
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
			input{n: 5, k: 2},
			output{ans: 3},
		},
		{
			input{n: 6, k: 5},
			output{ans: 1},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, findTheWinner(data.i.n, data.i.k), idx+1)
	}
}
