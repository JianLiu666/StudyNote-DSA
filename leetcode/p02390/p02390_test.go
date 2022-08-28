package p02390

import (
	"fmt"
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
			input{s: "leet**cod*e"},
			output{ans: "lecoe"},
		},
		{
			input{s: "erase*****"},
			output{ans: ""},
		},
		{
			input{s: "abb*cdfg*****x*"},
			output{ans: "a"},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, removeStars_bruteforce(data.i.s), fmt.Sprintf("Brute Force: %v", idx+1))
		ast.Equal(data.o.ans, removeStars_stack(data.i.s), fmt.Sprintf("Stack: %v", idx+1))
	}
}
