package p02309

// Time Complexity: O(n)
// Space Complexity: O(1)
func greatestLetter(s string) string {
	lowers := [26]int{}
	uppers := [26]int{}

	for _, ch := range s {
		if ch >= 'A' && ch <= 'Z' {
			uppers[int(ch)-65]++
		} else {
			lowers[int(ch)-97]++
		}
	}

	for i := len(uppers) - 1; i >= 0; i-- {
		if lowers[i] > 0 && uppers[i] > 0 {
			return string(byte(i + 65))
		}
	}

	return ""
}
