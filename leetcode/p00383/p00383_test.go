package p00383

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	ransomNote string
	magazine   string
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
			input{ransomNote: "a", magazine: "b"},
			output{ans: false},
		},
		{
			input{ransomNote: "aa", magazine: "ab"},
			output{ans: false},
		},
		{
			input{ransomNote: "aa", magazine: "aab"},
			output{ans: true},
		},
	}

	for _, data := range tds {
		ast.Equal(data.o.ans, canConstruct(data.i.ransomNote, data.i.magazine))
	}
}
