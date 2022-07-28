package p00494

// Time Complexity: O(n)
// Space Complexity: O(n)
func findTargetSumWays_dfs_memorization(nums []int, target int) int {
	memo := make([]map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		memo[i] = map[int]int{}
	}

	count := dfs_memorization(memo, nums, target, 0, 0, 0)
	return count
}

func dfs_memorization(memo []map[int]int, nums []int, target, idx, sum, count int) int {
	if idx == len(nums) {
		if sum == target {
			return 1
		}
		return 0
	}

	if val, exists := memo[idx][sum]; exists {
		return val
	}

	sum1 := dfs_memorization(memo, nums, target, idx+1, sum+nums[idx], count)
	sum2 := dfs_memorization(memo, nums, target, idx+1, sum-nums[idx], count)
	memo[idx][sum] = sum1 + sum2
	return sum1 + sum2
}
