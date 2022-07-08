package p01299

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	arr []int
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
			input{arr: []int{17, 18, 5, 4, 6, 1}},
			output{ans: []int{18, 6, 6, 6, 1, -1}},
		},
		{
			input{arr: []int{400}},
			output{ans: []int{-1}},
		},
	}

	for _, data := range tds {
		ast.Equal(data.o.ans, replaceElements(data.i.arr))
	}
}
