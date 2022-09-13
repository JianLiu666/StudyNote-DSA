package p00004

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	nums1 []int
	nums2 []int
}

type output struct {
	ans float64
}

type testdata struct {
	i input
	o output
}

func TestQuestion(t *testing.T) {
	ast := assert.New(t)

	tds := []testdata{
		{
			input{nums1: []int{1, 3}, nums2: []int{2}},
			output{ans: 2.00000},
		},
		{
			input{nums1: []int{1, 2}, nums2: []int{3, 4}},
			output{ans: 2.50000},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, findMedianSortedArrays(data.i.nums1, data.i.nums2), idx+1)
	}
}
