package p00189

// Time Complexity: O(n)
// Space Complexity: O(1)
func rotate(nums []int, k int) {
	k %= len(nums)

	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)
}

func reverse(nums []int, left, right int) {
	if left > right {
		return
	}

	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}

// Time Complexity: O(n)
// Space Complexity: O(1)
func rotate_cyclic(nums []int, k int) {
	k %= len(nums)
	start := 0
	count := 0

	for count < len(nums) {
		current := start
		prev := nums[start]

		for {
			next := (current + k) % len(nums)
			tmp := nums[next]
			nums[next] = prev
			prev = tmp
			current = next
			count++

			if current == start {
				break
			}
		}

		start++
	}
}
