package p00069

import "fmt"

// Time Complexity: O(logn)
// Space Complexity: O(1)
func mySqrt(x int) int {
	if x <= 1 {
		return x
	}

	head := 1
	tail := x

	for head < tail {
		mid := head + (tail-head)/2

		pow := mid * mid
		if pow == x {
			return mid
		} else if pow > x {
			tail = mid - 1
		} else {
			head = mid + 1
		}

		fmt.Println(head, tail)
	}

	if head*head > x {
		return head - 1
	}
	return head
}
