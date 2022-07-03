package q00268

// Time Complexity: O(n)
// Space Complexity: O(1)
func missingNumber_xor(nums []int) int {
	full_xor := 0
	nums_xor := 0

	for idx, num := range nums {
		full_xor ^= idx
		nums_xor ^= num
	}
	full_xor ^= len(nums)

	return full_xor ^ nums_xor
}

// Time Complexity: O(n)
// Space Complexity: O(1)
func missingNumber_math(nums []int) int {
	l := len(nums)
	actual_sum := l * (l + 1) / 2
	real_sum := 0

	for _, num := range nums {
		real_sum += num
	}

	return actual_sum - real_sum
}

// Time Complexity: O(n)
// Space Complexity: O(n)
func missingNumber_map(nums []int) int {
	seen := make(map[int]int, len(nums))

	for _, num := range nums {
		seen[num] = 1
	}

	for i := 0; i < len(nums); i++ {
		if _, exists := seen[i]; !exists {
			return i
		}
	}

	return len(nums)
}

// Time Complexity: O(nlogn)
// Space Complexity: O(n)
func missingNumber_sort(nums []int) int {
	sorted := merge_sort(nums)

	for idx, num := range nums {
		if idx != num {
			return idx
		}
	}

	return len(sorted)
}

func merge_sort(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}

	mid := len(nums) / 2
	l_arr := merge_sort(append([]int{}, nums[:mid]...))
	r_arr := merge_sort(append([]int{}, nums[mid:]...))

	l_idx := 0
	r_idx := 0
	current := 0

	for l_idx < len(l_arr) && r_idx < len(r_arr) {
		if l_arr[l_idx] <= r_arr[r_idx] {
			nums[current] = l_arr[l_idx]
			l_idx++
		} else {
			nums[current] = r_arr[r_idx]
			r_idx++
		}
		current++
	}

	for l_idx < len(l_arr) {
		nums[current] = l_arr[l_idx]
		l_idx++
		current++
	}

	for r_idx < len(r_arr) {
		nums[current] = r_arr[r_idx]
		r_idx++
		current++
	}

	return nums
}
