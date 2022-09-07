package p01249

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	s string
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
			input{s: "lee(t(c)o)de)"},
			output{ans: "lee(t(c)o)de"},
		},
		{
			input{s: "lee(t(code"},
			output{ans: "leetcode"},
		},
		{
			input{s: "()()"},
			output{ans: "()()"},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, minRemoveToMakeValid_index(data.i.s), fmt.Sprintf("Index: %d", idx+1))
		ast.Equal(data.o.ans, minRemoveToMakeValid_string(data.i.s), fmt.Sprintf("String: %d", idx+1))
	}
}
