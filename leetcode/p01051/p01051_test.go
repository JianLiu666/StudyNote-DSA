package p01051

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	heights []int
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
			input{heights: []int{1, 1, 4, 2, 1, 3}},
			output{ans: 3},
		},
		{
			input{heights: []int{5, 1, 2, 3, 4}},
			output{ans: 5},
		},
		{
			input{heights: []int{23, 52, 46, 7, 50, 87, 20, 32, 85, 65, 62, 34, 8, 86, 15, 66, 66, 30, 11, 96, 18, 26, 24, 10, 57, 13, 37, 69, 85, 6, 8, 17, 40, 88, 14, 72, 85, 51, 40, 38, 54, 65, 65, 27, 18, 59, 77, 12, 25, 46, 10, 19, 10, 28, 64, 79, 5, 88, 2, 1, 14, 50, 91, 34, 58, 32, 90, 67, 28, 81, 84, 76, 88, 45, 42, 54, 59, 56, 20, 6, 56, 51, 72, 69, 6, 48, 67, 68, 6, 10, 93, 69, 4, 29, 28}},
			output{ans: 95},
		},
	}

	for _, data := range tds {
		ast.Equal(data.o.ans, heightChecker(data.i.heights))
	}
}
