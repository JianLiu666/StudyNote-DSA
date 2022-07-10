package p00013

// Time Complexity: O(n)
// Space Complexity: O(1)
func romanToInt(s string) int {
	memo := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	ans := 0
	previous := 0
	for _, ch := range s {
		num := memo[ch]
		if num <= previous {
			ans += num
		} else {
			ans += num - previous*2
		}
		previous = num
	}

	return ans
}
