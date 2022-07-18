package p00118

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	rowIndex int
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
			input{rowIndex: 3},
			output{ans: []int{1, 3, 3, 1}},
		},
		{
			input{rowIndex: 0},
			output{ans: []int{1}},
		},
		{
			input{rowIndex: 1},
			output{ans: []int{1, 1}},
		},
	}

	for _, data := range tds {
		ast.Equal(data.o.ans, getRow(data.i.rowIndex))
	}
}
