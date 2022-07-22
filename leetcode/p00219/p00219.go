package p00219

// Time Complexity: O(n)
// Space Complexity: O(n)
func containsNearbyDuplicate(nums []int, k int) bool {
	if len(nums) <= 1 {
		return false
	}

	seen := make(map[int]int, len(nums))

	for idx, num := range nums {
		if idx2, exists := seen[num]; exists {
			if idx-idx2 <= k {
				return true
			}
		}
		seen[num] = idx
	}

	return false
}
