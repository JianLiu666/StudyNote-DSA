package p00438

// Time Complexity: O(nk), where n is the length of s, k is the length of p
// Space Complexity: O(k)
func findAnagrams(s string, p string) []int {
	result := []int{}

	sMap := map[rune]int{}
	pMap := map[rune]int{}
	for _, r := range p {
		pMap[r]++
	}

	for i, r := range s {
		sMap[r]++
		if i >= len(p) {
			sMap[rune(s[i-len(p)])]--
		}

		if i >= len(p)-1 && same(sMap, pMap) {
			result = append(result, i-len(p)+1)
		}
	}

	return result
}

func same(map1, map2 map[rune]int) bool {
	for key := range map1 {
		if map1[key] != map2[key] {
			return false
		}
	}
	return true
}
