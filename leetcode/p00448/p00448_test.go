package p00448

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	nums []int
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
			input{nums: []int{4, 3, 2, 7, 8, 2, 3, 1}},
			output{ans: []int{5, 6}},
		},
		{
			input{nums: []int{1, 1}},
			output{ans: []int{2}},
		},
	}

	for _, data := range tds {
		ast.Equal(data.o.ans, findDisappearedNumbers(data.i.nums))
	}
}
