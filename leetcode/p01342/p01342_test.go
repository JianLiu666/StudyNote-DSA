package p01342

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	num int
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
			input{num: 14},
			output{ans: 6},
		},
		{
			input{num: 8},
			output{ans: 4},
		},
		{
			input{num: 123},
			output{ans: 12},
		},
	}

	for _, data := range tds {
		ast.Equal(data.o.ans, numberOfSteps(data.i.num))
	}
}
