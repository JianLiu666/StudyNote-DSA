package p00209

// Time Complexity: O(n)
// Space Complexity: O(1)
func minSubArrayLen(target int, nums []int) int {
	head := 0
	sum := 0
	length := len(nums)

	for current := 0; current < len(nums); current++ {
		sum += nums[current]

		for sum >= target {
			if length > current-head+1 {
				length = current - head + 1
			}

			sum -= nums[head]
			head++
		}
	}

	if head == 0 && sum < target {
		return 0
	}
	return length
}
