package p00804

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	words []string
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
			input{words: []string{"gin", "zen", "gig", "msg"}},
			output{ans: 2},
		},
		{
			input{words: []string{"a"}},
			output{ans: 1},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, uniqueMorseRepresentations(data.i.words), idx+1)
	}
}
