package p02367

// Time Complexity: O(n)
// Space Complexity: O(n)
func arithmeticTriplets(nums []int, diff int) int {
	memo := map[int]bool{}
	for _, n := range nums {
		memo[n] = true
	}

	cnt := 0
	for _, n := range nums {
		if memo[n-diff] && memo[n+diff] {
			cnt++
		}
	}

	return cnt
}
