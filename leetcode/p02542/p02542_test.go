package p02542

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	nums1 []int
	nums2 []int
	k     int
}

type output struct {
	ans int64
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
				nums1: []int{1, 3, 3, 2},
				nums2: []int{2, 1, 3, 4},
				k:     3,
			},
			output{ans: 12},
		},
		{
			input{
				nums1: []int{4, 2, 3, 1, 1},
				nums2: []int{7, 5, 10, 9, 6},
				k:     1,
			},
			output{ans: 30},
		},
		{
			input{
				nums1: []int{23, 16, 20, 7, 3},
				nums2: []int{19, 21, 22, 22, 12},
				k:     3,
			},
			output{ans: 1121},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, maxScore(data.i.nums1, data.i.nums2, data.i.k), idx+1)
	}
}
