package p01578

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	colors     string
	neededTime []int
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
			input{colors: "abaac", neededTime: []int{1, 2, 3, 4, 5}},
			output{ans: 3},
		},
		{
			input{colors: "abc", neededTime: []int{1, 2, 3}},
			output{ans: 0},
		},
		{
			input{colors: "aabaa", neededTime: []int{1, 2, 3, 4, 1}},
			output{ans: 2},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, minCost(data.i.colors, data.i.neededTime), idx+1)
	}
}
