package p02007

import "sort"

// Time Complexity: O(nlogn), where n is the number of changed
// Space Complexity: O(n)
func findOriginalArray(changed []int) []int {
	if len(changed)&1 == 1 {
		return []int{}
	}

	res := []int{}
	sort.Ints(changed)

	memo := map[int]int{}
	for _, n := range changed {
		memo[n]++
	}

	for _, n := range changed {
		if memo[n] > 0 {
			memo[n*2]--
			if memo[n*2] < 0 {
				return []int{}
			}
			memo[n]--
			res = append(res, n)
		}
	}

	return res
}
