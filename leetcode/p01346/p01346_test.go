package p01346

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	arr []int
}

type output struct {
	ans bool
}

type testdata struct {
	i input
	o output
}

func TestQuestion(t *testing.T) {
	ast := assert.New(t)

	tds := []testdata{
		{
			input{arr: []int{10, 2, 5, 3}},
			output{ans: true},
		},
		{
			input{arr: []int{7, 1, 14, 11}},
			output{ans: true},
		},
		{
			input{arr: []int{3, 1, 7, 11}},
			output{ans: false},
		},
	}

	for _, data := range tds {
		ast.Equal(data.o.ans, checkIfExist(data.i.arr))
	}
}
