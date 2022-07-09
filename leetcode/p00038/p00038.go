package p00038

import (
	"fmt"
	"strings"
)

// Time Complexity: O(n^2)
// Space Complexity: O(n)
func countAndSay_byte(n int) string {
	if n == 1 {
		return "1"
	}

	tmp := countAndSay_byte(n - 1)

	var res strings.Builder
	keep := tmp[0]
	count := 1
	for i := 1; i < len(tmp); i++ {
		if tmp[i] != keep {
			res.WriteByte(byte(count + 48))
			res.WriteByte(keep)
			keep = tmp[i]
			count = 1
		} else {
			count++
		}
	}
	res.WriteByte(byte(count + 48))
	res.WriteByte(keep)

	return res.String()
}

// Time Complexity: O(n^2)
// Space Complexity: O(n)
func countAndSay_bruteforce(n int) string {
	if n == 1 {
		return "1"
	}

	tmp := countAndSay_bruteforce(n - 1)

	res := ""
	keep := tmp[0]
	count := 1
	for i := 1; i < len(tmp); i++ {
		if tmp[i] != keep {
			res = fmt.Sprintf("%s%d%c", res, count, keep)
			keep = tmp[i]
			count = 1
		} else {
			count++
		}
	}
	res = fmt.Sprintf("%s%d%c", res, count, keep)

	return res
}
