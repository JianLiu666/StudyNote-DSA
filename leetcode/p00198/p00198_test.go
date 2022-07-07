package p00198

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
			input{nums: []int{1, 2, 3, 1}},
			output{ans: 4},
		},
		{
			input{nums: []int{2, 7, 9, 3, 1}},
			output{ans: 12},
		},
	}

	for _, data := range tds {
		ast.Equal(data.o.ans, rob(data.i.nums))
	}
}
