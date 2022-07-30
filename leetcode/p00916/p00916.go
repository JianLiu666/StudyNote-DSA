package p00916

// Time Complexity: O(n+k), where k is the length of words2
// Space Complexity: O(1)
func wordSubsets(words1 []string, words2 []string) []string {
	subsetMap := [26]int{}
	for _, word := range words2 {
		letters := counter(word)
		for i := 0; i < 26; i++ {
			if letters[i] > subsetMap[i] {
				subsetMap[i] = letters[i]
			}
		}
	}

	res := []string{}
	for _, word := range words1 {
		supersetMap := counter(word)
		isSubset := true
		for i := 0; i < 26; i++ {
			if subsetMap[i] > supersetMap[i] {
				isSubset = false
				break
			}
		}
		if isSubset {
			res = append(res, word)
		}
	}

	return res
}

func counter(word string) [26]int {
	frequency := [26]int{}
	size := len(word)
	for i := 0; i < size; i++ {
		frequency[word[i]-'a']++
	}
	return frequency
}
