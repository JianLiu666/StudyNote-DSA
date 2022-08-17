package p01925

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
			input{n: 5},
			output{ans: 2},
		},
		{
			input{n: 10},
			output{ans: 4},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, countTriples(data.i.n), idx+1)
	}
}
