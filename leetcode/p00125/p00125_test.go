package p00125

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	s string
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
			input{s: "A man, a plan, a canal: Panama"},
			output{ans: true},
		},
		{
			input{s: "race a car"},
			output{ans: false},
		},
		{
			input{s: " "},
			output{ans: true},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, isPalindrome(data.i.s), idx+1)
	}
}
