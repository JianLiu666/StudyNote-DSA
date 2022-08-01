package p00278

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
			input{n: 5},
			output{ans: 4},
		},
		{
			input{n: 1},
			output{ans: 1},
		},
		{
			input{n: 2126753390},
			output{ans: 1702766719},
		},
	}

	for _, data := range tds {
		badVersion = data.o.ans
		ast.Equal(data.o.ans, firstBadVersion(data.i.n))
	}
}
