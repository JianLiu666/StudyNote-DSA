package p02389

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	nums    []int
	queries []int
}

type output struct {
	ans []int
}

type testdata struct {
	i input
	o output
}

func TestQuestion(t *testing.T) {
	ast := assert.New(t)

	tds := []testdata{
		{
			input{nums: []int{4, 5, 2, 1}, queries: []int{3, 10, 21}},
			output{ans: []int{2, 3, 4}},
		},
		{
			input{nums: []int{2, 3, 4, 5}, queries: []int{1}},
			output{ans: []int{0}},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, answerQueries(data.i.nums, data.i.queries), idx+1)
	}
}
