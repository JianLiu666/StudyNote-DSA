package p00162

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
			output{ans: 2},
		},
		{
			input{nums: []int{1, 2, 1, 3, 5, 6, 4}},
			output{ans: 5},
		},
	}

	for _, data := range tds {
		ast.Equal(data.o.ans, findPeakElement(data.i.nums))
	}
}
