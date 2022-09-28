package p00461

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	x int
	y int
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
			input{x: 1, y: 4},
			output{ans: 2},
		},
		{
			input{x: 3, y: 1},
			output{ans: 1},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, hammingDistance(data.i.x, data.i.y), idx+1)
	}
}
