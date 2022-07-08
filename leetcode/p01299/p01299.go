package p01299

// Time Complexity: O(n)
// Space Complexity: O(1)
func replaceElements(arr []int) []int {
	tail := len(arr) - 1
	maximum := arr[tail]
	arr[tail] = -1

	for i := tail - 1; i > -1; i-- {
		if arr[i] > maximum {
			arr[i], maximum = maximum, arr[i]
		} else {
			arr[i] = maximum
		}
	}

	return arr
}
