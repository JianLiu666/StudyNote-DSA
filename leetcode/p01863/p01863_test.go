package p01863

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
			input{nums: []int{1, 3}},
			output{ans: 6},
		},
		{
			input{nums: []int{5, 1, 6}},
			output{ans: 28},
		},
		{
			input{nums: []int{3, 4, 5, 6, 7, 8}},
			output{ans: 480},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, subsetXORSum(data.i.nums), idx+1)
	}
}
