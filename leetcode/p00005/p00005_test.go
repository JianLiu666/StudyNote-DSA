package p00005

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
			input{s: "babad"},
			output{ans: "bab"},
		},
		{
			input{s: "cbbd"},
			output{ans: "bb"},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, longestPalindrome_bruteforce(data.i.s), fmt.Sprintf("Brute Force: %v", idx+1))
		ast.Equal(data.o.ans, longestPalindrome_bruteforce_optimize(data.i.s), fmt.Sprintf("Brute Force Optimize: %v", idx+1))
	}
}
