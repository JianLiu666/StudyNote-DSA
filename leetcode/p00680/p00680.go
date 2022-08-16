package p00680

// Time Complexity: O(n)
// Space Complexity: O(1)
func validPalindrome(s string) bool {
	head, tail := 0, len(s)-1
	for head < tail {
		if s[head] != s[tail] {
			return validSubString(s[head:tail]) || validSubString(s[head+1:tail+1])
		}
		head++
		tail--
	}

	return true
}

func validSubString(s string) bool {
	head, tail := 0, len(s)-1
	for head < tail {
		if s[head] != s[tail] {
			return false
		}
		head++
		tail--
	}

	return true
}
