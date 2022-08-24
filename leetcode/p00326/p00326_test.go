package p00326

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	n int
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
			input{n: 27},
			output{ans: true},
		},
		{
			input{n: 0},
			output{ans: false},
		},
		{
			input{n: 9},
			output{ans: true},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, isPowerOfThree_recursion(data.i.n), fmt.Sprintf("Recursion: %v", idx+1))
		ast.Equal(data.o.ans, isPowerOfThree_limitation(data.i.n), fmt.Sprintf("Integer Limitation: %v", idx+1))
	}
}
