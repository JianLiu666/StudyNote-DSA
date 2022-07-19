package p00557

import (
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
			input{s: "Let's take LeetCode contest"},
			output{ans: "s'teL ekat edoCteeL tsetnoc"},
		},
		{
			input{s: "God Ding"},
			output{ans: "doG gniD"},
		},
	}

	for _, data := range tds {
		ast.Equal(data.o.ans, reverseWords(data.i.s))
	}
}
