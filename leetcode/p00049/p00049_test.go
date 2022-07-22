package p00049

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	strs []string
}

type output struct {
	ans [][]string
}

type testdata struct {
	i input
	o output
}

func TestQuestion(t *testing.T) {
	ast := assert.New(t)

	tds := []testdata{
		{
			input{strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"}},
			output{ans: [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}}},
		},
		{
			input{strs: []string{""}},
			output{ans: [][]string{{""}}},
		},
		{
			input{strs: []string{"a"}},
			output{ans: [][]string{{"a"}}},
		},
	}

	for _, data := range tds {
		ast.Equal(data.o.ans, groupAnagrams(data.i.strs))
	}
}
