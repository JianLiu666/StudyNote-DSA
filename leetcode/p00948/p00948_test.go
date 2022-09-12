package p00948

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	tokens []int
	power  int
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
			input{tokens: []int{100}, power: 50},
			output{ans: 0},
		},
		{
			input{tokens: []int{100, 200}, power: 150},
			output{ans: 1},
		},
		{
			input{tokens: []int{100, 200, 300, 400}, power: 200},
			output{ans: 2},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, bagOfTokensScore(data.i.tokens, data.i.power), idx+1)
	}
}
