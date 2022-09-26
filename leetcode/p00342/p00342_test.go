package p00342

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
			input{n: 16},
			output{ans: true},
		},
		{
			input{n: 5},
			output{ans: false},
		},
		{
			input{n: 1},
			output{ans: true},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, isPowerOfFour_recursion(data.i.n), fmt.Sprintf("Recursion: %v", idx+1))
		ast.Equal(data.o.ans, isPowerOfFour_bitmanipulation(data.i.n), fmt.Sprintf("Bit Manipulation: %v", idx+1))
	}
}
