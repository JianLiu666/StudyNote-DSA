package p00494

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
			input{nums: []int{1, 1, 1, 1, 1}, target: 3},
			output{ans: 5},
		},
		{
			input{nums: []int{1, 1, 1, 1, 1, 1, 1, 1, 1}, target: 3},
			output{ans: 5},
		},
		{
			input{nums: []int{1}, target: 1},
			output{ans: 1},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, findTargetSumWays_dfs_bruteforce(data.i.nums, data.i.target), fmt.Sprintf("dfs_bruteforce: %d", idx+1))
		ast.Equal(data.o.ans, findTargetSumWays_dfs_memorization(data.i.nums, data.i.target), fmt.Sprintf("dfs_memorization: %d", idx+1))
	}
}
