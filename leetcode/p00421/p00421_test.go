package p00421

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
			input{nums: []int{3, 10, 5, 25, 2, 8}},
			output{ans: 28},
		},
		{
			input{nums: []int{14, 70, 53, 83, 49, 91, 36, 80, 92, 51, 66, 70}},
			output{ans: 127},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, findMaximumXOR(data.i.nums), idx+1)
	}
}
