package p00007

import (
	"math"
	"strconv"
)

// Time Complexity: O(n)
// Space Complexity: O(n)
func reverse_string(x int) int {
	s := strconv.Itoa(x)

	head := 0
	if s[head] == '-' {
		head++
	}

	r := ""
	for i := len(s) - 1; i >= head; i-- {
		r += string(s[i])
	}

	num, _ := strconv.Atoi(r)
	if num > math.MaxInt32-1 {
		return 0
	}

	if x > 0 {
		return num
	}
	return num * -1
}

// Time Complexity: O(n)
// Space Complexity: O(1)
func reverse_recursion(x int) int {
	sig := 1
	if x < 0 {
		x *= -1
		sig = -1
	}

	res := 0
	for x != 0 {
		num := x % 10
		res = res*10 + num
		x /= 10
	}

	if res > math.MaxInt32-1 {
		return 0
	}

	return res * sig
}
