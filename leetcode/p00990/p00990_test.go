package p00990

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	equations []string
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
			input{equations: []string{"a==b", "b!=a"}},
			output{ans: false},
		},
		{
			input{equations: []string{"b==a", "a==b"}},
			output{ans: true},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, equationsPossible(data.i.equations), idx+1)
	}
}
