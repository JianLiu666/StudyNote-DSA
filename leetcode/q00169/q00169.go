package q00169

// Boyer-Moore Voting Algorithm
// Time Complexity: O(n)
// Space Complexity: O(1)
func majorityElement(nums []int) int {
	max_num := 0
	count := 0
	for _, num := range nums {
		if count == 0 {
			max_num = num
			count = 1
		} else if max_num == num {
			count++
		} else {
			count--
		}
	}

	return max_num
}

// Divide and Conquer
// Time Complexity: O(nlogn)
// Space Complexity: O(logn)
func majorityElement_dac(nums []int) int {
	return recursive(nums, 0, len(nums)-1)
}

func recursive(nums []int, head, tail int) int {
	if head == tail {
		return nums[head]
	}

	mid := (tail-head)/2 + head
	left := recursive(nums, head, mid)
	right := recursive(nums, mid+1, tail)

	if left == right {
		return left
	}

	left_count := 0
	right_count := 0
	for i := head; i <= tail; i++ {
		if nums[i] == left {
			left_count++
		} else if nums[i] == right {
			right_count++
		}
	}

	if left_count > right_count {
		return left
	}
	return right
}

// Time Complexity: O(n)
// Space Complexity: O(n)
func majorityElement_map(nums []int) int {
	seen := make(map[int]int, len(nums)/2+1)

	max_value := 0
	count := 0
	for _, num := range nums {
		seen[num] += 1
		if seen[num] > count {
			max_value = num
			count = seen[num]
		}
	}

	return max_value
}
