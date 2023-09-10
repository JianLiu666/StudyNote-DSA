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
	count := 0

	memo := map[int]int{}
	memo[0] = 1 // 表示 preSum = 0 的時候出現了 1 次

	preSum := 0
	for i := 0; i < len(nums); i++ {
		preSum += nums[i]

		// 這題題目的問題可以用 preSum[j] - preSum[i-1] == k 來表示 subarray[i,j] == k
		// 所以我們可以求 preSum[j] - k == preSum[i-1] 的次數
		// 來反推 preSum[j] - preSum[i-1] == k
		count += memo[preSum-k]

		memo[preSum]++
	}

	return count
}
