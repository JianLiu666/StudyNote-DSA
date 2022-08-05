package p00377

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	nums   []int
	target int
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
			input{nums: []int{1, 2, 3}, target: 4},
			output{ans: 7},
		},
		{
			input{nums: []int{9}, target: 3},
			output{ans: 0},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, combinationSum4_dpTopDown(data.i.nums, data.i.target), fmt.Sprintf("DP Top-Down: %v", idx+1))
		ast.Equal(data.o.ans, combinationSum4_dpBottomUp(data.i.nums, data.i.target), fmt.Sprintf("DP Bottom-Up: %v", idx+1))
	}
}
