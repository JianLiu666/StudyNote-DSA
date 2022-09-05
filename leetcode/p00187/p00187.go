package p00187

// Time Complexity: O(n)
// Space Complexity: O(n)
func findRepeatedDnaSequences(s string) []string {
	memo := map[string]int{}
	res := []string{}

	for i := 0; i <= len(s)-10; i++ {
		memo[s[i:i+10]]++
		if memo[s[i:i+10]] == 2 {
			res = append(res, s[i:i+10])
		}
	}

	return res
}
