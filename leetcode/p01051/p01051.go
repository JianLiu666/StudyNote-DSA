package p01051

// Time Complexity: O(n)
// Space Complexity: O(n)
func heightChecker(heights []int) int {
	counter := make([]int, len(heights)+1)

	// Time COmplexity: O(n)
	for _, v := range heights {
		counter[v-1]++
	}

	mismatch := 0
	current := 0

	// Time COmplexity: O(n)
	for i, v := range counter {
		for j := 0; j < v; j++ {
			if heights[current] != i+1 {
				mismatch++
			}
			current++
		}
	}

	return mismatch
}
