package p00049

import "strings"

// Time Complexity: O(n^2)
// Space Complexity: O(n)
func groupAnagrams(strs []string) [][]string {
	res := make([][]string, 0)

	anagrams_group := map[string]int{}
	group_index := 0

	for _, str := range strs {
		key := hasher(str)
		if idx, exists := anagrams_group[key]; exists {
			res[idx] = append(res[idx], str)
		} else {
			res = append(res, []string{str})
			anagrams_group[key] = group_index
			group_index++
		}
	}

	return res
}

// Time Complexity: O(n)
// Space Complexity: O(1)
func hasher(str string) string {
	var res strings.Builder

	mapping := [26]int{}
	for _, ch := range str {
		mapping[ch-97]++
	}

	for idx, count := range mapping {
		for i := 0; i < count; i++ {
			res.WriteByte(byte(idx + 97))
		}
	}

	return res.String()
}
