package p00134

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	gas  []int
	cost []int
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
			input{
				gas:  []int{1, 2, 3, 4, 5},
				cost: []int{3, 4, 5, 1, 2},
			},
			output{
				ans: 3,
			},
		},
		{
			input{
				gas:  []int{2, 3, 4},
				cost: []int{3, 4, 3},
			},
			output{
				ans: -1,
			},
		},
		{
			input{
				gas:  []int{5, 1, 2, 3, 4},
				cost: []int{4, 4, 1, 5, 1},
			},
			output{
				ans: 4,
			},
		},
		{
			input{
				gas:  []int{3, 1, 1},
				cost: []int{1, 2, 2},
			},
			output{
				ans: 0,
			},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, canCompleteCircuit(data.i.gas, data.i.cost), idx+1)
	}
}
