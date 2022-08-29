package p00238

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	nums []int
}

type output struct {
	ans []int
}

type testdata struct {
	i input
	o output
}

func TestQuestion(t *testing.T) {
	ast := assert.New(t)

	tds := []testdata{
		{
			input{nums: []int{1, 2, 3, 4}},
			output{ans: []int{24, 12, 8, 6}},
		},
		{
			input{nums: []int{-1, 1, 0, -3, 3}},
			output{ans: []int{0, 0, 9, 0, 0}},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, productExceptSelf_bruteforce(data.i.nums), fmt.Sprintf("Brute Force: %v", idx+1))
		ast.Equal(data.o.ans, productExceptSelf_prefixsum(data.i.nums), fmt.Sprintf("Prefix Sum: %v", idx+1))
	}
}
