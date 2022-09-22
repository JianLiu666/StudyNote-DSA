package p00839

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	strs []string
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
			input{strs: []string{"tars", "rats", "arts", "star"}},
			output{ans: 2},
		},
		{
			input{strs: []string{"omv", "ovm"}},
			output{ans: 1},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, numSimilarGroups(data.i.strs), idx+1)
	}
}
