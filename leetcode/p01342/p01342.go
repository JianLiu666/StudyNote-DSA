package p01342

// Time Complexity: O(logn)
// Space Complexity: O(1)
func numberOfSteps(num int) int {
	if num == 0 {
		return 0
	}

	count := -1
	for num != 0 {
		if num&1 == 0 {
			count += 1
		} else {
			count += 2
		}

		num >>= 1
	}

	return count
}
