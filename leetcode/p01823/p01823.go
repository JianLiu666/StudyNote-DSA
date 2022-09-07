package p01823

// Time Complexity: O(n)
// Space Complexity: O(n)
func findTheWinner(n int, k int) int {
	queue := make([]int, n)
	for i := 0; i < n; i++ {
		queue[i] = i + 1
	}

	cur := 0
	for len(queue) > 1 {
		cur = (cur + k - 1) % len(queue)
		queue = append(queue[:cur], queue[cur+1:]...)
	}

	return queue[0]
}
