package q00121

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	prices []int
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
			input{prices: []int{7, 1, 5, 3, 6, 4}},
			output{ans: 5},
		},
		{
			input{prices: []int{7, 6, 4, 3, 1}},
			output{ans: 0},
		},
	}

	for _, data := range tds {
		ast.Equal(data.o.ans, maxProfit(data.i.prices))
	}
}
