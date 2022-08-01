package p00658

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	arr []int
	k   int
	x   int
}

type output struct {
	ans []int
}

type testdata struct {
	i input
	o output
}

func TestQuestion(t *testing.T) {
	ast := assert.New(t)

	tds := []testdata{
		{
			input{arr: []int{1, 2, 3, 4, 5}, k: 4, x: 3},
			output{ans: []int{1, 2, 3, 4}},
		},
		{
			input{arr: []int{1, 2, 3, 4, 5}, k: 4, x: -1},
			output{ans: []int{1, 2, 3, 4}},
		},
	}

	for _, data := range tds {
		ast.Equal(data.o.ans, findClosestElements(data.i.arr, data.i.k, data.i.x))
	}
}
