package p00334

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	nums []int
}

type output struct {
	ans bool
}

type testdata struct {
	i input
	o output
}

func TestQuestion(t *testing.T) {
	ast := assert.New(t)

	tds := []testdata{
		{
			input{nums: []int{1, 2, 3, 4, 5}},
			output{ans: true},
		},
		{
			input{nums: []int{5, 4, 3, 2, 1}},
			output{ans: false},
		},
		{
			input{nums: []int{2, 1, 5, 0, 4, 6}},
			output{ans: true},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, increasingTriplet(data.i.nums), idx+1)
	}
}
