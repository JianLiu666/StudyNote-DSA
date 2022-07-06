package p00509

import "math"

// Time Complexity: O(2^n)
// Space Complexity: O(n)
func fib_recursion(n int) int {
	if n < 2 {
		return n
	}

	return fib_recursion(n-1) + fib_recursion(n-2)
}

// Time Complexity: O(n)
// Space Complexity: O(1)
func fib_dp(n int) int {
	if n < 2 {
		return n
	}

	fn, fn_1, fn_2 := 0, 1, 0
	for i := 1; i < n; i++ {
		fn = fn_1 + fn_2
		fn_2 = fn_1
		fn_1 = fn
	}

	return fn
}

// Time Complexity: O(1)
// Space Complexity: O(1)
func fib_math(n int) int {
	phi := (math.Sqrt(5) + 1) / 2

	return int(math.Round(math.Pow(phi, float64(n)) / math.Sqrt(5)))
}
