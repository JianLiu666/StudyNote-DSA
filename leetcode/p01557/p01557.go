package p01557

// Time Complexity: O(n)
// Space Complexity: O(n)
func findSmallestSetOfVertices(n int, edges [][]int) []int {
	indegrees := make([]int, n)
	for _, edge := range edges {
		indegrees[edge[1]]++
	}

	res := []int{}
	for idx, count := range indegrees {
		if count == 0 {
			res = append(res, idx)
		}
	}

	return res
}
