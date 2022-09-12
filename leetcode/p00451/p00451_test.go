package p00451

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	s string
}

type output struct {
	ans string
}

type testdata struct {
	i input
	o output
}

func TestQuestion(t *testing.T) {
	ast := assert.New(t)

	tds := []testdata{
		{
			input{s: "acccddddbb"},
			output{ans: "ddddcccbba"},
		},
	}

	for _, data := range tds {
		ast.Equal(data.o.ans, frequencySort(data.i.s))
	}
}
