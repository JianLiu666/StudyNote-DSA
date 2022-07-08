package p00374

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

var answer = 0

// Time Complexity: O(logn)
// Space Complexity: O(1)
func guessNumber(n int) int {
	head := 1
	tail := n

	for true {
		num := head + (tail-head)/2

		res := guess(num)
		if res == 0 {
			return num
		} else if res == 1 {
			head = num + 1
		} else {
			tail = num - 1
		}
	}

	return -1
}

func guess(n int) int {
	if n == answer {
		return 0
	} else if n < answer {
		return 1
	}

	return -1
}
