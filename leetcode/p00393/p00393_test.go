package p00393

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	data []int
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
			input{data: []int{115, 100, 102, 231, 154, 132, 13, 10}},
			output{ans: true},
		},
		{
			input{data: []int{197, 130, 1}},
			output{ans: true},
		},
		{
			input{data: []int{235, 140, 4}},
			output{ans: false},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, validUtf8(data.i.data), idx+1)
	}
}
