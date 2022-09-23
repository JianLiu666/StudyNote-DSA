package p01200

import (
	"math"
	"sort"
)

// Time Complexity: O(nlogn)
// Space Complexity: O(1)
func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)

	res := [][]int{}
	min := math.MaxInt
	for i := 1; i < len(arr); i++ {
		diff := arr[i] - arr[i-1]
		if diff < min {
			res = [][]int{{arr[i-1], arr[i]}}
			min = diff
		} else if diff == min {
			res = append(res, []int{arr[i-1], arr[i]})
		}
	}

	return res
}
