package p00941

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
			input{arr: []int{2, 1}},
			output{ans: false},
		},
		{
			input{arr: []int{3, 5, 5}},
			output{ans: false},
		},
		{
			input{arr: []int{0, 3, 2, 1}},
			output{ans: true},
		},
	}

	for _, data := range tds {
		ast.Equal(data.o.ans, validMountainArray(data.i.arr))
	}
}
