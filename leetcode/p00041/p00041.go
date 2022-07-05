package p00041

// Time Complexity: O(n)
// Space Complexity: O(1)
func firstMissingPositive_swap(nums []int) int {
	for i := range nums {
		for nums[i] >= 1 && nums[i] <= len(nums) && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}

	for i := range nums {
		if i+1 != nums[i] {
			return i + 1
		}
	}

	return len(nums) + 1
}

// Time Complexity: O(n)
// Space Complexity: O(1)
func firstMissingPositive_hash(nums []int) int {
	nums = append(nums, 0)

	size := len(nums)
	for i := range nums {
		if nums[i] < 0 || nums[i] > size {
			nums[i] = 0
		}
	}

	for i := range nums {
		nums[nums[i]%size] += size
	}

	for i := 1; i < size; i++ {
		if nums[i]/size == 0 {
			return i
		}
	}

	return size
}
