package p00202

// Time Complexity: O(logn)
// Space Complexity: O(logn)
func isHappy(n int) bool {
	seen := map[int]bool{}
	for {
		n = getNext(n)
		if n == 1 {
			return true
		}
		if _, exists := seen[n]; exists {
			return false
		} else {
			seen[n] = true
		}
	}
}

// Time Complexity: O(logn)
// Space Complexity: O(1)
func isHappy_twoPointers(n int) bool {
	slow := n
	fast := n
	for fast != 1 && getNext(fast) != 1 {
		slow = getNext(slow)
		fast = getNext(getNext(fast))

		if slow == fast {
			return false
		}
	}

	return true
}

// Time Complexity: O(logn)
// Space Complexity: O(1)
func isHappy_math(n int) bool {
	for {
		if n == 1 {
			return true
		} else if n == 4 {
			return false
		}

		n = getNext(n)
	}
}

// Time Complexity: O(logn)
// Space Complexity: O(1)
func getNext(n int) int {
	table := map[int]int{
		0: 0,
		1: 1,
		2: 4,
		3: 9,
		4: 16,
		5: 25,
		6: 36,
		7: 49,
		8: 64,
		9: 81,
	}

	sum := 0
	for n != 0 {
		remainder := n % 10
		sum += table[remainder]
		n /= 10
	}

	return sum
}
