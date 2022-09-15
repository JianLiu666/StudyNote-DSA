package p02007

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	changed []int
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
			input{changed: []int{1, 3, 4, 2, 6, 8}},
			output{ans: []int{1, 3, 4}},
		},
		{
			input{changed: []int{6, 3, 0, 1}},
			output{ans: []int{}},
		},
		{
			input{changed: []int{1}},
			output{ans: []int{}},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, findOriginalArray(data.i.changed), idx+1)
	}
}
