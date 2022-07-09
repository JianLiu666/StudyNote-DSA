package p00038

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	n int
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
			input{n: 1},
			output{ans: "1"},
		},
		{
			input{n: 2},
			output{ans: "11"},
		},
		{
			input{n: 3},
			output{ans: "21"},
		},
		{
			input{n: 4},
			output{ans: "1211"},
		},
	}

	for _, data := range tds {
		ast.Equal(data.o.ans, countAndSay_byte(data.i.n))
		// ast.Equal(data.o.ans, countAndSay_bruteforce(data.i.n))
	}
}
