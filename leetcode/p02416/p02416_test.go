package p02416

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	words []string
}

type output struct {
	ans []int
}

type testdata struct {
	i input
	o output
}

func TestQuestion(t *testing.T) {
	ast := assert.New(t)

	tds := []testdata{
		{
			input{words: []string{"abc", "ab", "bc", "b"}},
			output{ans: []int{5, 4, 3, 2}},
		},
		{
			input{words: []string{"abcd"}},
			output{ans: []int{4}},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, sumPrefixScores(data.i.words), idx+1)
	}
}
