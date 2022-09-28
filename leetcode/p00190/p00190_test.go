package p00190

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	num uint32
}

type output struct {
	ans uint32
}

type testdata struct {
	i input
	o output
}

func TestQuestion(t *testing.T) {
	ast := assert.New(t)

	tds := []testdata{
		{
			input{num: 0b00000010100101000001111010011100},
			output{ans: 0b00111001011110000010100101000000},
		},
		{
			input{num: 0b11111111111111111111111111111101},
			output{ans: 0b10111111111111111111111111111111},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, reverseBits(data.i.num), idx+1)
	}
}
