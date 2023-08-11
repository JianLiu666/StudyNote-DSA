package p00049

// Time Complexity: O(n^2)
// Space Complexity: O(n)
func groupAnagrams(strs []string) [][]string {
	result := [][]string{}

	groups := map[[26]int][]string{}

	for _, s := range strs {
		// 透過 array 解決 string immutable 的問題
		key := [26]int{}
		for _, r := range s {
			key[r-'a']++
		}
		groups[key] = append(groups[key], s)
	}

	for _, group := range groups {
		result = append(result, group)
	}

	return result
}
