package p00048

import (
	"fmt"
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
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}},
			output{ans: [][]int{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3},
			}},
		},
		{
			input{matrix: [][]int{
				{5, 1, 9, 11},
				{2, 4, 8, 10},
				{13, 3, 6, 7},
				{15, 14, 12, 16},
			}},
			output{ans: [][]int{
				{15, 13, 2, 5},
				{14, 3, 4, 1},
				{12, 6, 8, 9},
				{16, 7, 10, 11},
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
		rotate_linearalgebra(copyright)
		ast.Equal(data.o.ans, copyright, fmt.Sprintf("Linear Algebra: %v", idx+1))

		copyright = make([][]int, len(data.i.matrix))
		for i := 0; i < len(data.i.matrix); i++ {
			arr := make([]int, len(data.i.matrix[i]))
			copy(arr, data.i.matrix[i])
			copyright[i] = arr
		}
		rotate_bruteforce(copyright)
		ast.Equal(data.o.ans, copyright, fmt.Sprintf("Brute Force: %v", idx+1))
	}
}
