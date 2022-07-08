package p01346

// Time Complexity: O(n)
// Space Complexity: O(n)
func checkIfExist(arr []int) bool {
	memo := make(map[int]int, len(arr))

	for _, n := range arr {
		if _, exists := memo[n]; exists {
			return true
		}

		memo[n*2] = 1
		if n%2 == 0 {
			memo[n/2] = 1
		}
	}

	return false
}
