package p00050

import "math"

// Time Complexity: O(logn)
// Space Complexity: O(1)
func myPow(x float64, n int) float64 {
	if x == 0 {
		return 0
	}
	if n == 0 {
		return 1
	}

	neg := false
	if n < 0 {
		n = -n
		neg = true
	}

	var ans float64 = 1
	for n != 0 {
		if n%2 != 0 {
			ans *= x
		}
		x *= x
		n >>= 1
	}

	if neg {
		ans = 1 / ans
	}
	return math.Round(ans*100000) / 100000
}
