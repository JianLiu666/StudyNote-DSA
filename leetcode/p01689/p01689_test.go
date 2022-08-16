package p01689

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	s string
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
			input{s: "32"},
			output{ans: 3},
		},
		{
			input{s: "82734"},
			output{ans: 8},
		},
		{
			input{s: "27346209830709182346"},
			output{ans: 9},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, minPartitions(data.i.s), idx+1)
	}
}
