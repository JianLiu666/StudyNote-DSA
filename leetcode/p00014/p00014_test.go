package p00014

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	strs []string
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
			input{strs: []string{"flower", "flow", "flight"}},
			output{ans: "fl"},
		},
		{
			input{strs: []string{"dog", "racecar", "car"}},
			output{ans: ""},
		},
	}

	for _, data := range tds {
		ast.Equal(data.o.ans, longestCommonPrefix(data.i.strs))
	}
}
