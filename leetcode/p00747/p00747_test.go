package p00747

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
			input{nums: []int{3, 6, 1, 0}},
			output{ans: 1},
		},
		{
			input{nums: []int{1, 2, 3, 4}},
			output{ans: -1},
		},
	}

	for _, data := range tds {
		ast.Equal(data.o.ans, dominantIndex(data.i.nums))
	}
}
