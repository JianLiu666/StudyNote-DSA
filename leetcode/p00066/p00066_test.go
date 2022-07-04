package p00066

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	digits []int
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
			input{digits: []int{1, 2, 3}},
			output{ans: []int{1, 2, 4}},
		},
		{
			input{digits: []int{4, 3, 2, 1}},
			output{ans: []int{4, 3, 2, 2}},
		},
		{
			input{digits: []int{9}},
			output{ans: []int{1, 0}},
		},
	}

	for _, data := range tds {
		ast.Equal(data.o.ans, plusOne(data.i.digits))
	}
}
