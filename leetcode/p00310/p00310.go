package p00310

// Time Complexity: O(v+e), where v is the number of nodes, e is the number of edges
// Space Complexity: O(n)
func findMinHeightTrees(n int, edges [][]int) []int {
	if len(edges) == 0 {
		return []int{0}
	}

	numInDegrees := make([]int, n)
	outDegrees := make([][]int, n)
	for i := 0; i < len(edges); i++ {
		numInDegrees[edges[i][0]]++
		outDegrees[edges[i][0]] = append(outDegrees[edges[i][0]], edges[i][1])
		numInDegrees[edges[i][1]]++
		outDegrees[edges[i][1]] = append(outDegrees[edges[i][1]], edges[i][0])

	}

	q := []int{}
	for i, n := range numInDegrees {
		if n == 1 {
			q = append(q, i)
		}
	}

	for n > 2 {
		size := len(q)
		n -= size
		for ; size > 0; size-- {
			idx := q[0]
			q = q[1:]
			numInDegrees[idx]--
			for _, next := range outDegrees[idx] {
				numInDegrees[next]--
				if numInDegrees[next] == 1 {
					q = append(q, next)
				}
			}
		}
	}

	return q
}
