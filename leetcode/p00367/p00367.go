package p00367

// Time Complexity: O(logn)
// Space Complexity: O(1)
func isPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}

	low := 1
	high := num

	for low <= high {
		mid := low + (high-low)/2
		square := mid * mid
		if square == num {
			return true
		} else if square > num {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return false
}
