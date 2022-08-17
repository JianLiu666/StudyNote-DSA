package p01925

// Time Complexity: O(n^2)
// Space Complexity: O(n)
func countTriples(n int) int {
	memo := map[int]bool{}
	for i := 1; i <= n; i++ {
		memo[i*i] = true
	}

	cnt := 0
	for i := 1; i <= n; i++ {
		for j := i + 1; j <= n; j++ {
			if memo[i*i+j*j] {
				cnt += 2
			}
		}
	}

	return cnt
}
