package p00053

// Time Complexity: O(n)
// Space Complexity: O(1)
func maxSubArray(nums []int) int {
	res := nums[0]
	cur := nums[0]

	for i := 1; i < len(nums); i++ {
		if cur < 0 || cur+nums[i] < 0 {
			cur = nums[i]
		} else {
			cur += nums[i]
		}

		if cur > res {
			res = cur
		}
	}

	return res
}
