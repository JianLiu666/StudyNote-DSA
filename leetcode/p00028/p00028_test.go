package p00028

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	haystack string
	needle   string
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
			input{haystack: "hello", needle: "ll"},
			output{ans: 2},
		},
		{
			input{haystack: "aaaaa", needle: "bba"},
			output{ans: -1},
		},
		{
			input{haystack: "mississippi", needle: "issip"},
			output{ans: 4},
		},
	}

	for _, data := range tds {
		ast.Equal(data.o.ans, strStr(data.i.haystack, data.i.needle))
	}
}
