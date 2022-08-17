package p02367

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	nums []int
	diff int
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
			input{nums: []int{0, 1, 4, 6, 7, 10}, diff: 3},
			output{ans: 2},
		},
		{
			input{nums: []int{4, 5, 6, 7, 8, 9}, diff: 2},
			output{ans: 2},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, arithmeticTriplets(data.i.nums, data.i.diff), idx+1)
	}
}
