package p00377

// Time Complexity: O(nm), where n is the length of nums, and m is the number of target
// Space Complexity: O(m), where n is the number of target
func combinationSum4_dpBottomUp(nums []int, target int) int {
	memo := make([]int, target+1)
	memo[0] = 1
	for i := 1; i <= target; i++ {
		for _, n := range nums {
			if i-n >= 0 {
				memo[i] += memo[i-n]
			}
		}
	}

	return memo[target]
}

// Time Complexity: O(n), where n is the number of target
// Space Complexity: O(n), where n is the number of target
func combinationSum4_dpTopDown(nums []int, target int) int {
	memo := make([]int, target+1)
	memo[0] = 1
	for i := 1; i < len(memo); i++ {
		memo[i] = -1
	}

	dpTopDown(memo, nums, target)
	return memo[target]
}

func dpTopDown(memo, nums []int, target int) int {
	if memo[target] != -1 {
		return memo[target]
	}

	res := 0
	for i := 0; i < len(nums); i++ {
		if target >= nums[i] {
			res += dpTopDown(memo, nums, target-nums[i])
		}
	}

	memo[target] = res
	return res
}
