package p00560

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	nums []int
	k    int
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
			input{nums: []int{1, 1, 1}, k: 2},
			output{ans: 2},
		},
		{
			input{nums: []int{1, 2, 3}, k: 3},
			output{ans: 2},
		},
		{
			input{nums: []int{-1, -1, 1}, k: 0},
			output{ans: 1},
		},
		{
			input{nums: []int{1}, k: 0},
			output{ans: 0},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, subarraySum_bruteforce(data.i.nums, data.i.k), fmt.Sprintf("Brute Force: %v", idx+1))
		ast.Equal(data.o.ans, subarraySum_presum(data.i.nums, data.i.k), fmt.Sprintf("Prefix Sum: %v", idx+1))
	}
}
