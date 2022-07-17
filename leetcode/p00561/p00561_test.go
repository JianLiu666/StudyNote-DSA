package p00561

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
			input{nums: []int{1, 4, 3, 2}},
			output{ans: 4},
		},
		{
			input{nums: []int{6, 2, 6, 5, 1, 2}},
			output{ans: 9},
		},
	}

	for _, data := range tds {
		ast.Equal(data.o.ans, arrayPairSum(data.i.nums))
	}
}
