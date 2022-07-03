package q00283

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
			input{nums: []int{0, 1, 0, 3, 12}},
			output{ans: []int{1, 3, 12, 0, 0}},
		},
		{
			input{nums: []int{0}},
			output{ans: []int{0}},
		},
	}

	for _, data := range tds {
		moveZeroes(data.i.nums)
		ast.Equal(data.o.ans, data.i.nums)
	}
}
