package p00718

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
				nums1: []int{1, 2, 3, 2, 1},
				nums2: []int{3, 2, 1, 4, 7},
			},
			output{ans: 3},
		},
		{
			input{
				nums1: []int{0, 0, 0, 0, 0},
				nums2: []int{0, 0, 0, 0, 0},
			},
			output{ans: 5},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, findLength(data.i.nums1, data.i.nums2), idx+1)
	}
}
