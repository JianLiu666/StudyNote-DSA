package q00088

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	nums1 []int
	m     int
	nums2 []int
	n     int
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
			input{nums1: []int{1, 2, 3, 0, 0, 0}, m: 3, nums2: []int{2, 5, 6}, n: 3},
			output{ans: []int{1, 2, 2, 3, 5, 6}},
		},
		{
			input{nums1: []int{1}, m: 1, nums2: []int{}, n: 0},
			output{ans: []int{1}},
		},
		{
			input{nums1: []int{0}, m: 0, nums2: []int{1}, n: 1},
			output{ans: []int{1}},
		},
	}

	for _, data := range tds {
		merge(data.i.nums1, data.i.m, data.i.nums2, data.i.n)
		ast.Equal(data.o.ans, data.i.nums1)
	}
}
