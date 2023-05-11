package p01035

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	nums1 []int
	nums2 []int
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
			input{
				nums1: []int{1, 4, 2},
				nums2: []int{1, 2, 4},
			},
			output{
				ans: 2,
			},
		},
		{
			input{
				nums1: []int{2, 5, 1, 2, 5},
				nums2: []int{10, 5, 2, 1, 5, 2},
			},
			output{
				ans: 3,
			},
		},
		{
			input{
				nums1: []int{1, 3, 7, 1, 7, 5},
				nums2: []int{1, 9, 2, 5, 1},
			},
			output{
				ans: 2,
			},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, maxUncrossedLines(data.i.nums1, data.i.nums2), idx+1)
	}
}
