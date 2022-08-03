package p00744

// Time Complexity: O(logn)
// Space Complexity: O(1)
func nextGreatestLetter(letters []byte, target byte) byte {
	if letters[len(letters)-1] <= target {
		return letters[0]
	}

	low, high := 0, len(letters)-1

	for low < high {
		mid := low + (high-low)/2
		if letters[mid] <= target {
			low = mid + 1
		} else {
			high = mid
		}
	}

	return letters[low]
}
