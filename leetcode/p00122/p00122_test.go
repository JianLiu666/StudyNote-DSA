package p00122

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
			output{ans: 7},
		},
		{
			input{prices: []int{1, 2, 3, 4, 5}},
			output{ans: 4},
		},
		{
			input{prices: []int{7, 6, 4, 3, 1}},
			output{ans: 0},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, maxProfit(data.i.prices), idx+1)
	}
}
