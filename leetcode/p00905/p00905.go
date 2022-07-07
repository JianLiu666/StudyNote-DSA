package p00905

// Time Complexity: O(n)
// Space Complexity: O(1)
func sortArrayByParity(nums []int) []int {
	head := 0
	tail := len(nums) - 1

	for head < tail {
		for nums[head]%2 == 0 && head < tail {
			head++
		}
		for nums[tail]%2 == 1 && head < tail {
			tail--
		}

		if head < tail {
			nums[head], nums[tail] = nums[tail], nums[head]
			head++
			tail--
		}
	}

	return nums
}
