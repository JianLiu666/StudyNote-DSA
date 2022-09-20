package p00210

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	numCourses    int
	prerequisites [][]int
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
			input{
				numCourses:    2,
				prerequisites: [][]int{{1, 0}},
			},
			output{
				ans: []int{0, 1},
			},
		},
		{
			input{
				numCourses:    4,
				prerequisites: [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}},
			},
			output{
				ans: []int{0, 1, 2, 3},
			},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, findOrder(data.i.numCourses, data.i.prerequisites), idx+1)
	}
}
