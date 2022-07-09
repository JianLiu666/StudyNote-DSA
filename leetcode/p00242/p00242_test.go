package p00242

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	s string
	t string
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
			input{s: "anagram", t: "nagaram"},
			output{ans: true},
		},
		{
			input{s: "rat", t: "car"},
			output{ans: false},
		},
	}

	for _, data := range tds {
		ast.Equal(data.o.ans, isAnagram(data.i.s, data.i.t))
	}
}
