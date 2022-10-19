package p00692

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	words []string
	k     int
}

type output struct {
	ans []string
}

type testdata struct {
	i input
	o output
}

func TestQuestion(t *testing.T) {
	ast := assert.New(t)

	tds := []testdata{
		{
			input{words: []string{"i", "love", "leetcode", "i", "love", "coding"}, k: 2},
			output{ans: []string{"i", "love"}},
		},
		{
			input{words: []string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}, k: 4},
			output{ans: []string{"the", "is", "sunny", "day"}},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, topKFrequent(data.i.words, data.i.k), idx+1)
	}
}
