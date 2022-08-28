package p02389

import "sort"

// Time Complexity: O(mn), where m is the length of nums, n is the length of queries
// Space Complexity: O(1)
func answerQueries(nums []int, queries []int) []int {
	sort.Ints(nums)

	res := make([]int, len(queries))
	for i, query := range queries {
		for _, num := range nums {
			if query-num >= 0 {
				query -= num
				res[i]++
			} else {
				break
			}
		}
	}

	return res
}
