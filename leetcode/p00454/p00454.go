package p00454

// Time Complexity: O(n^2)
// Space Complexity: O(n^2)
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	lookup := map[int]int{}

	for _, n1 := range nums1 {
		for _, n2 := range nums2 {
			lookup[n1+n2]++
		}
	}

	count := 0
	for _, n3 := range nums3 {
		for _, n4 := range nums4 {
			count += lookup[-(n3 + n4)]
		}
	}

	return count
}
