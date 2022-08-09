package p00823

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	arr []int
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
			input{arr: []int{2, 4}},
			output{ans: 3},
		},
		{
			input{arr: []int{2, 4, 5, 10}},
			output{ans: 7},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, numFactoredBinaryTrees(data.i.arr), idx+1)
	}
}
