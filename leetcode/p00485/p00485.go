package p00485

// Time Complexity: O(n)
// Space Complexity: O(1)
func findMaxConsecutiveOnes(nums []int) int {
	maximum := 0
	counter := 0

	for _, num := range nums {
		if num == 1 {
			counter++
			if counter > maximum {
				maximum = counter
			}
		} else {
			counter = 0
		}
	}

	return maximum
}
