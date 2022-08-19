package sorting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeSort(t *testing.T) {
	testdata := [][]int{
		{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		{5, 5, 5, 5, 5, 5},
		{1, 1, 2, 2, 4, 4, 5, 5},
		{5, 4, 3, 2, 2, 1, 2, 2, 4, 3},
	}
	groundtruth := [][]int{
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		{5, 5, 5, 5, 5, 5},
		{1, 1, 2, 2, 4, 4, 5, 5},
		{1, 2, 2, 2, 2, 3, 3, 4, 4, 5},
	}

	ast := assert.New(t)
	for idx, nums := range testdata {
		mergesort(nums)
		ast.Equal(groundtruth[idx], nums)
	}
}
