package p00030

// Time Complexity: O(nm), where n is num of words, and m is length of each word
// Space Complexity: O(n), where n is num of words
func findSubstring(s string, words []string) []int {
	memo := map[string]int{}
	for _, word := range words {
		memo[word]++
	}

	res := []int{}
	skip := len(words[0])
	for start := 0; start < skip; start++ {
		anchor := start
		tmp := map[string]int{}
		count := 0
		for i := start; i <= len(s)-skip; i += skip {
			word := s[i : i+skip]
			if _, exists := memo[word]; exists {
				if tmp[word]+1 <= memo[word] {
					tmp[word]++
					count++
					if count == len(words) {
						res = append(res, anchor)
						count--
						tmp[s[anchor:anchor+skip]]--
						anchor += skip
					}
				} else {
					previous := s[anchor : anchor+skip]
					for previous != word {
						tmp[previous]--
						count--
						anchor += skip
						previous = s[anchor : anchor+skip]
					}
					anchor += skip
				}
			} else {
				tmp = map[string]int{}
				count = 0
				anchor = i + skip
			}
		}
	}

	return res
}
