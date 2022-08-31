package p00763

// Time Complexity: O(n)
// Space Complexity: O(n)
func partitionLabels(s string) []int {
	memo := map[byte]int{}
	for i := 0; i < len(s); i++ {
		memo[s[i]] = i
	}

	res := []int{}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		if memo[s[i]] > end {
			end = memo[s[i]]
		}
		if i == end {
			res = append(res, end-start+1)
			start = i + 1
		}
	}

	return res
}
