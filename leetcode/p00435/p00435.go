package p00435

import "sort"

// Time Complexity: O(nlogn)
// Space Complexity: O(1)
func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] <= intervals[j][1]
	})

	end := intervals[0][1]
	removed := 0
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= end {
			end = intervals[i][1]
		} else {
			removed++
		}
	}

	return removed
}
