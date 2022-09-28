package p00055

import (
	"fmt"
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
			input{nums: []int{2, 3, 1, 1, 4}},
			output{ans: true},
		},
		{
			input{nums: []int{3, 2, 1, 0, 4}},
			output{ans: false},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, canJump_bottomUp(data.i.nums), fmt.Sprintf("Bottom Up: %v", idx+1))
		ast.Equal(data.o.ans, canJump_bottomUp_optimize(data.i.nums), fmt.Sprintf("Bottom Up Optimize: %v", idx+1))
	}
}
