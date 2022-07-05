package p00027

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	nums []int
	val  int
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
			input{nums: []int{3, 2, 2, 3}, val: 3},
			output{ans: 2},
		},
		{
			input{nums: []int{0, 1, 2, 2, 3, 0, 4, 2}, val: 2},
			output{ans: 5},
		},
	}

	for _, data := range tds {
		ast.Equal(data.o.ans, removeElement(data.i.nums, data.i.val))
	}
}
