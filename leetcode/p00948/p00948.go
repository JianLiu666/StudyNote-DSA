package p00948

import "sort"

// Time Complexity: O(nlogn), where n is the number of tokens
// Space Complexity: O(1)
func bagOfTokensScore(tokens []int, power int) int {
	sort.Ints(tokens)

	current, maximum := 0, 0
	head, tail := 0, len(tokens)-1
	for head <= tail {
		if power >= tokens[head] {
			power -= tokens[head]
			head++
			current++
			if current > maximum {
				maximum = current
			}

		} else if current > 0 {
			power += tokens[tail]
			tail--
			current--

		} else {
			break
		}
	}

	return maximum
}
