package p00387

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	s string
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
			input{s: "leetcode"},
			output{ans: 0},
		},
		{
			input{s: "loveleetcode"},
			output{ans: 2},
		},
		{
			input{s: "aabb"},
			output{ans: -1},
		},
	}

	for _, data := range tds {
		ast.Equal(data.o.ans, firstUniqChar(data.i.s))
	}
}
