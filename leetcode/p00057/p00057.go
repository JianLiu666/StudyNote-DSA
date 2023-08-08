package p00057

// Time Complexity: O(n)
// Space Complexity: O(n)
func insert(intervals [][]int, newInterval []int) [][]int {
	result := [][]int{}
	cursor := 0

	// 把所有小於 newInterval.start 的都先加進 results
	for cursor < len(intervals) && intervals[cursor][1] < newInterval[0] {
		result = append(result, intervals[cursor])
		cursor++
	}

	// 開始合併介於 newInterval.start、newInterval.end 中間的 intervals
	// 特別注意 intervals[i].start 剛好等於 newInterval.start 時也需要合併
	for cursor < len(intervals) && intervals[cursor][0] <= newInterval[1] {
		newInterval[0] = min(newInterval[0], intervals[cursor][0])
		newInterval[1] = max(newInterval[1], intervals[cursor][1])
		cursor++
	}
	result = append(result, newInterval)

	// 把剩下的 intervals 加進 results
	for cursor < len(intervals) {
		result = append(result, intervals[cursor])
		cursor++
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
