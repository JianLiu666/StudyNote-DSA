package p00322

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	coins  []int
	amount int
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
			input{coins: []int{1, 2, 5}, amount: 11},
			output{ans: 3},
		},
		{
			input{coins: []int{2}, amount: 3},
			output{ans: -1},
		},
		{
			input{coins: []int{1}, amount: 0},
			output{ans: 0},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, coinChange(data.i.coins, data.i.amount), idx+1)
	}
}
