package p00240

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	matrix [][]int
	target int
}

type output struct {
	ans bool
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
				matrix: [][]int{
					{1, 4, 7, 11, 15},
					{2, 5, 8, 12, 19},
					{3, 6, 9, 16, 22},
					{10, 13, 14, 17, 24},
					{18, 21, 23, 26, 30}},
				target: 5},
			output{ans: true},
		},
		{
			input{
				matrix: [][]int{
					{1, 4, 7, 11, 15},
					{2, 5, 8, 12, 19},
					{3, 6, 9, 16, 22},
					{10, 13, 14, 17, 24},
					{18, 21, 23, 26, 30}},
				target: 20},
			output{ans: false},
		},
		{
			input{
				matrix: [][]int{
					{-1},
					{-1}},
				target: 0},
			output{ans: false},
		},
		{
			input{
				matrix: [][]int{
					{1, 2, 3, 4, 5},
					{6, 7, 8, 9, 10},
					{11, 12, 13, 14, 15},
					{16, 17, 18, 19, 20},
					{21, 22, 23, 24, 25}},
				target: 5},
			output{ans: true},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, searchMatrix(data.i.matrix, data.i.target), idx+1)
	}
}
