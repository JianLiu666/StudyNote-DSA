package p00152

func maxProduct(nums []int) int {
	result := nums[0]

	minDp := make([]int, len(nums))
	maxDp := make([]int, len(nums))
	minDp[0] = nums[0]
	maxDp[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		maxDp[i] = max(max(maxDp[i-1]*nums[i], minDp[i-1]*nums[i]), nums[i])
		minDp[i] = min(min(maxDp[i-1]*nums[i], minDp[i-1]*nums[i]), nums[i])
		result = max(result, maxDp[i])
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
