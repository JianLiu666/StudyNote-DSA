package p00171

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	columnTitle string
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
			input{columnTitle: "A"},
			output{ans: 1},
		},
		{
			input{columnTitle: "AB"},
			output{ans: 28},
		},
		{
			input{columnTitle: "ZY"},
			output{ans: 701},
		},
	}

	for _, data := range tds {
		ast.Equal(data.o.ans, titleToNumber(data.i.columnTitle))
	}
}
