package p00455

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	g []int
	s []int
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
				g: []int{1, 2, 3},
				s: []int{1, 1},
			},
			output{ans: 1},
		},
		{
			input{
				g: []int{1, 2},
				s: []int{1, 2, 3},
			},
			output{ans: 2},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, findContentChildren(data.i.g, data.i.s), idx+1)
	}
}
