package q00217

// Time: O(n)
// Space: O(n)
func containsDuplicate_map(nums []int) bool {
	seen := make(map[int]int, len(nums))

	for _, num := range nums {
		if _, exists := seen[num]; exists {
			return true
		}

		seen[num] = 1
	}

	return false
}

// Time: O(nlogn)
// Space: O(n)
func containsDuplicate_sort(nums []int) bool {
	sorted := merge_sort(nums)
	for i := 1; i < len(sorted); i++ {
		if sorted[i-1] == sorted[i] {
			return true
		}
	}

	return false
}

// Merge sort
// Top-down recursion
func merge_sort(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}

	mid := len(nums) / 2
	left_arr := merge_sort(append([]int{}, nums[:mid]...))
	right_arr := merge_sort(append([]int{}, nums[mid:]...))

	left_idx := 0
	right_idx := 0
	current_idx := 0

	for left_idx < len(left_arr) && right_idx < len(right_arr) {
		if left_arr[left_idx] <= right_arr[right_idx] {
			nums[current_idx] = left_arr[left_idx]
			left_idx++
		} else {
			nums[current_idx] = right_arr[right_idx]
			right_idx++
		}
		current_idx++
	}

	for left_idx < len(left_arr) {
		nums[current_idx] = left_arr[left_idx]
		left_idx++
		current_idx++
	}

	for right_idx < len(right_arr) {
		nums[current_idx] = right_arr[right_idx]
		right_idx++
		current_idx++
	}

	return nums
}
