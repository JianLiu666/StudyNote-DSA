package p00560

// Time Complexity: O(n^2)
// Space Complexity: O(1)
func subarraySum_bruteforce(nums []int, k int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum == k {
				count++
			}
		}
	}

	return count
}

// Time Complexity: O(n)
// Space Complexity: O(n)
func subarraySum_presum(nums []int, k int) int {
	memo := map[int]int{0: 1}
	preSum := 0
	count := 0

	for i := 0; i < len(nums); i++ {
		preSum += nums[i]

		if memo[preSum-k] > 0 {
			count += memo[preSum-k]
		}
		memo[preSum]++
	}

	return count
}
