package p00179

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	nums []int
}

type output struct {
	ans string
}

type testdata struct {
	i input
	o output
}

func TestQuestion(t *testing.T) {
	ast := assert.New(t)

	tds := []testdata{
		{
			input{nums: []int{10, 2}},
			output{ans: "210"},
		},
		{
			input{nums: []int{3, 30, 34, 5, 9}},
			output{ans: "9534330"},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, largestNumber(data.i.nums), idx+1)
	}
}
