package p00164

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
			input{nums: []int{3, 6, 9, 1}},
			output{ans: 3},
		},
		{
			input{nums: []int{10}},
			output{ans: 0},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, maximumGap(data.i.nums), idx+1)
	}
}
