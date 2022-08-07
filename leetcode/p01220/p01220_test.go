package p01220

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	n int
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
			input{n: 1},
			output{ans: 5},
		},
		{
			input{n: 2},
			output{ans: 10},
		},
		{
			input{n: 5},
			output{ans: 68},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, countVowelPermutation_optimize(data.i.n), fmt.Sprintf("optimize: %d", idx+1))
		ast.Equal(data.o.ans, countVowelPermutation_array(data.i.n), fmt.Sprintf("array: %d", idx+1))
	}
}
