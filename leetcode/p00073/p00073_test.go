package p00073

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	matrix [][]int
}

type output struct {
	ans [][]int
}

type testdata struct {
	i input
	o output
}

func TestQuestion(t *testing.T) {
	ast := assert.New(t)

	tds := []testdata{
		{
			input{matrix: [][]int{
				{1, 1, 1},
				{1, 0, 1},
				{1, 1, 1},
			}},
			output{ans: [][]int{
				{1, 0, 1},
				{0, 0, 0},
				{1, 0, 1},
			}},
		},
		{
			input{matrix: [][]int{
				{0, 1, 2, 0},
				{3, 4, 5, 2},
				{1, 3, 1, 5},
			}},
			output{ans: [][]int{
				{0, 0, 0, 0},
				{0, 4, 5, 0},
				{0, 3, 1, 0},
			}},
		},
	}

	for idx, data := range tds {
		copyright := make([][]int, len(data.i.matrix))
		for i := 0; i < len(data.i.matrix); i++ {
			arr := make([]int, len(data.i.matrix[i]))
			copy(arr, data.i.matrix[i])
			copyright[i] = arr
		}
		setZeroes(copyright)
		ast.Equal(data.o.ans, copyright, idx+1)
	}
}
