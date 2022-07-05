package p00027

// Time Complexity: O(n)
// Space Complexity: O(1)
func removeElement(nums []int, val int) int {
	ptr := 0
	for i := range nums {
		if nums[i] != val {
			nums[ptr] = nums[i]
			ptr++
		}
	}

	return ptr
}
