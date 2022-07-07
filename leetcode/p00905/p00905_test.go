package p00905

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
			input{nums: []int{3, 1, 2, 4}},
			output{ans: []int{4, 2, 1, 3}},
		},
		{
			input{nums: []int{0}},
			output{ans: []int{0}},
		},
	}

	for _, data := range tds {
		ast.Equal(data.o.ans, sortArrayByParity(data.i.nums))

	}
}
