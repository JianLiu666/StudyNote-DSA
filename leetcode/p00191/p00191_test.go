package p00191

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	n uint32
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
			input{n: 11},
			output{ans: 3},
		},
		{
			input{n: 128},
			output{ans: 1},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, hammingWeight(data.i.n), idx+1)
	}
}
