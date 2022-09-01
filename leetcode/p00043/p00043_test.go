package p00043

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	num1 string
	num2 string
}

type output struct {
	ans string
}

type testdata struct {
	i input
	o output
}

func TestQuestion(t *testing.T) {
	ast := assert.New(t)

	tds := []testdata{
		{
			input{num1: "2", num2: "3"},
			output{ans: "6"},
		},
		{
			input{num1: "123", num2: "456"},
			output{ans: "56088"},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, multiply(data.i.num1, data.i.num2), idx+1)
	}
}
