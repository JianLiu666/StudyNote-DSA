package p00062

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	m int
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
			input{m: 3, n: 7},
			output{ans: 28},
		},
		{
			input{m: 3, n: 2},
			output{ans: 3},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, uniquePaths(data.i.m, data.i.n), idx+1)
	}
}
