package p00771

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	jewels string
	stones string
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
			input{jewels: "aA", stones: "aAAbbbb"},
			output{ans: 3},
		},
		{
			input{jewels: "z", stones: "ZZ"},
			output{ans: 0},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, numJewelsInStones(data.i.jewels, data.i.stones), idx+1)
	}
}
