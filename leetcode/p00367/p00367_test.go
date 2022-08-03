package p00367

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	num int
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
			input{num: 16},
			output{ans: true},
		},
		{
			input{num: 14},
			output{ans: false},
		},
		{
			input{num: 104976},
			output{ans: true},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, isPerfectSquare(data.i.num), idx+1)
	}
}
