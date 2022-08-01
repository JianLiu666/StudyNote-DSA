package p00542

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	mat [][]int
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
			input{mat: [][]int{
				{0, 0, 0},
				{0, 1, 0},
				{0, 0, 0}}},
			output{ans: [][]int{
				{0, 0, 0},
				{0, 1, 0},
				{0, 0, 0}}},
		},
		{
			input{mat: [][]int{
				{0, 0, 0},
				{0, 1, 0},
				{1, 1, 1}}},
			output{ans: [][]int{
				{0, 0, 0},
				{0, 1, 0},
				{1, 2, 1}}},
		},
		{
			input{mat: [][]int{
				{0, 1, 0},
				{0, 1, 0},
				{0, 1, 0},
				{0, 1, 0},
				{0, 1, 0}}},
			output{ans: [][]int{
				{0, 1, 0},
				{0, 1, 0},
				{0, 1, 0},
				{0, 1, 0},
				{0, 1, 0}}},
		},
		{
			input{mat: [][]int{
				{1, 1, 0, 0, 1, 0, 0, 1, 1, 0},
				{1, 0, 0, 1, 0, 1, 1, 1, 1, 1},
				{1, 1, 1, 0, 0, 1, 1, 1, 1, 0},
				{0, 1, 1, 1, 0, 1, 1, 1, 1, 1},
				{0, 0, 1, 1, 1, 1, 1, 1, 1, 0},
				{1, 1, 1, 1, 1, 1, 0, 1, 1, 1},
				{0, 1, 1, 1, 1, 1, 1, 0, 0, 1},
				{1, 1, 1, 1, 1, 0, 0, 1, 1, 1},
				{0, 1, 0, 1, 1, 0, 1, 1, 1, 1},
				{1, 1, 1, 0, 1, 0, 1, 1, 1, 1}}},
			output{ans: [][]int{
				{2, 1, 0, 0, 1, 0, 0, 1, 1, 0},
				{1, 0, 0, 1, 0, 1, 1, 2, 2, 1},
				{1, 1, 1, 0, 0, 1, 2, 2, 1, 0},
				{0, 1, 2, 1, 0, 1, 2, 3, 2, 1},
				{0, 0, 1, 2, 1, 2, 1, 2, 1, 0},
				{1, 1, 2, 3, 2, 1, 0, 1, 1, 1},
				{0, 1, 2, 3, 2, 1, 1, 0, 0, 1},
				{1, 2, 1, 2, 1, 0, 0, 1, 1, 2},
				{0, 1, 0, 1, 1, 0, 1, 2, 2, 3},
				{1, 2, 1, 0, 1, 0, 1, 2, 3, 4}}},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, updateMatrix_bfs(data.i.mat), fmt.Sprintf("dfs: %v", idx+1))
		ast.Equal(data.o.ans, updateMatrix_dp(data.i.mat), fmt.Sprintf("dp: %v", idx+1))
	}
}
