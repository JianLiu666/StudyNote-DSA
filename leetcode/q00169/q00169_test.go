package q00169

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	nums []int
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
			input{nums: []int{3, 2, 3}},
			output{ans: 3},
		},
		{
			input{nums: []int{2, 2, 1, 1, 1, 2, 2}},
			output{ans: 2},
		},
		{
			input{nums: []int{1}},
			output{ans: 1},
		},
	}

	for _, data := range tds {
		ast.Equal(data.o.ans, majorityElement(data.i.nums))
		ast.Equal(data.o.ans, majorityElement_dac(data.i.nums))
		ast.Equal(data.o.ans, majorityElement_map(data.i.nums))
	}
}
