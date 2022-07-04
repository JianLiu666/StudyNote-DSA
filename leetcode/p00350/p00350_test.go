package p00350

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	nums1 []int
	nums2 []int
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
			input{nums1: []int{1, 2, 2, 1}, nums2: []int{2, 2}},
			output{ans: []int{2, 2}},
		},
		{
			input{nums1: []int{4, 9, 5}, nums2: []int{9, 4, 9, 8, 4}},
			output{ans: []int{4, 9}},
		},
	}

	for _, data := range tds {
		sorted := intersect_dict(data.i.nums1, data.i.nums2)
		sort.Ints(sorted)
		ast.Equal(data.o.ans, sorted)

		sorted = intersect_sort(data.i.nums1, data.i.nums2)
		sort.Ints(sorted)
		ast.Equal(data.o.ans, sorted)
	}
}
