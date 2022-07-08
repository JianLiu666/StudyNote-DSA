package p00374

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	n    int
	pick int
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
			input{n: 10, pick: 6},
			output{ans: 6},
		},
		{
			input{n: 1, pick: 1},
			output{ans: 1},
		},
		{
			input{n: 2, pick: 1},
			output{ans: 1},
		},
	}

	for _, data := range tds {
		answer = data.i.pick
		ast.Equal(data.o.ans, guessNumber(data.i.n))
	}
}
