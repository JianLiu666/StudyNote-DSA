package p00494

// Time Complexity: O(2^n)
// Space Complexity: O(n)
func findTargetSumWays_dfs_bruteforce(nums []int, target int) int {
	return dfs_bruteforce(nums, target, 0, 0, 0)
}

func dfs_bruteforce(nums []int, target, idx, sum, count int) int {
	if idx == len(nums) {
		if sum == target {
			return 1
		}
		return 0
	}

	sum1 := dfs_bruteforce(nums, target, idx+1, sum+nums[idx], count)
	sum2 := dfs_bruteforce(nums, target, idx+1, sum-nums[idx], count)

	return sum1 + sum2
}
