package q00011

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	height []int
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
			input{height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7}},
			output{ans: 49},
		},
		{
			input{height: []int{1, 1}},
			output{ans: 1},
		},
	}

	for _, data := range tds {
		ast.Equal(data.o.ans, maxArea(data.i.height))
	}
}
