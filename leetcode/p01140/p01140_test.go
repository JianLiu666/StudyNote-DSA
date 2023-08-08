package p01140

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	piles []int
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
			input{piles: []int{2, 7, 9, 4, 4}},
			output{ans: 10},
		},
		{
			input{piles: []int{1, 2, 3, 4, 5, 100}},
			output{ans: 104},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, stoneGameII(data.i.piles), idx+1)
	}
}
