package p00646

import "sort"

// Time Complexiyt: O(nlogn)
// Space Complexity: O(1)
func findLongestChain(pairs [][]int) int {
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i][1] < pairs[j][1] {
			return true
		}
		if pairs[i][1] == pairs[j][1] {
			return pairs[i][0] <= pairs[j][0]
		}
		return false
	})

	count := 1
	cur := 0
	for i := 1; i < len(pairs); i++ {
		if pairs[cur][1] < pairs[i][0] {
			count++
			cur = i
		}
	}

	return count
}
