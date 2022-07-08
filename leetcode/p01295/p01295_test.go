package p01295

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
			input{nums: []int{12, 345, 2, 6, 7896}},
			output{ans: 2},
		},
		{
			input{nums: []int{555, 901, 482, 1771}},
			output{ans: 1},
		},
	}

	for _, data := range tds {
		ast.Equal(data.o.ans, findNumbers(data.i.nums))
	}
}
