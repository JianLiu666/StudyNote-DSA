package p00070

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	n int
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
			input{n: 2},
			output{ans: 2},
		},
		{
			input{n: 3},
			output{ans: 3},
		},
	}

	for _, data := range tds {
		ast.Equal(data.o.ans, climbStairs(data.i.n))
	}
}
