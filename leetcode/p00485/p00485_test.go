package p00485

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
			input{nums: []int{1, 1, 0, 1, 1, 1}},
			output{ans: 3},
		},
		{
			input{nums: []int{1, 0, 1, 1, 0, 1}},
			output{ans: 2},
		},
	}

	for _, data := range tds {
		ast.Equal(data.o.ans, findMaxConsecutiveOnes(data.i.nums))
	}
}
