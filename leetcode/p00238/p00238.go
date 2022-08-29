package p00238

// Time Complexity: O(n)
// Space Complexity: O(1)
func productExceptSelf_bruteforce(nums []int) []int {
	zero_count := 0
	product := 1 // without zero

	for _, n := range nums {
		if n != 0 {
			product *= n
		} else {
			zero_count++
		}
	}

	res := make([]int, len(nums))
	if zero_count > 1 {
		return res
	}

	for i := 0; i < len(nums); i++ {
		if zero_count == 0 {
			res[i] = product / nums[i]
		} else {
			if nums[i] == 0 {
				res[i] = product
			} else {
				res[i] = 0
			}
		}
	}

	return res
}

// Time Complexity: O(n)
// Space Complexity: O(1)
func productExceptSelf_prefixsum(nums []int) []int {
	res := make([]int, len(nums))
	res[0] = 1

	for i := 1; i < len(nums); i++ {
		res[i] = res[i-1] * nums[i-1]
	}

	right := 1
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] *= right
		right *= nums[i]
	}

	return res
}
