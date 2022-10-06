package p00567

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	s1 string
	s2 string
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
			input{s1: "ab", s2: "eidbaooo"},
			output{ans: true},
		},
		{
			input{s1: "ab", s2: "eidboaoo"},
			output{ans: false},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, checkInclusion(data.i.s1, data.i.s2), idx+1)
	}
}
