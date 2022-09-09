package p01996

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	properties [][]int
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
			input{properties: [][]int{{5, 5}, {6, 3}, {3, 6}}},
			output{ans: 0},
		},
		{
			input{properties: [][]int{{2, 2}, {3, 3}}},
			output{ans: 1},
		},
		{
			input{properties: [][]int{{1, 5}, {10, 4}, {4, 3}}},
			output{ans: 1},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, numberOfWeakCharacters(data.i.properties), idx+1)
	}
}
