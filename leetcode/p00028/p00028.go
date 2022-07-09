package p00028

// Time Complexity: O(n)
// Space Complexity: O(1)
func strStr(haystack string, needle string) int {
	haystack_size := len(haystack)
	needle_size := len(needle)

	if needle_size == 0 {
		return 0
	}
	if haystack_size < needle_size {
		return -1
	}

	for i := 0; i < haystack_size; i++ {
		if haystack[i] == needle[0] {
			if i+needle_size <= haystack_size {
				if haystack[i:i+needle_size] == needle {
					return i
				}
			} else {
				return -1
			}
		}
	}

	return -1
}
