package p00207

func canFinish_bfs(numCourses int, prerequisites [][]int) bool {

	numIndegrees := make([]int, numCourses) // 紀錄節點的入邊數量
	outDegrees := make([][]int, numCourses) // 紀錄節點能走訪的下個節點位置

	// 掃一遍 graph 把 numInDegrees, outDegrees 都記錄好
	for _, pair := range prerequisites {
		numIndegrees[pair[0]]++
		outDegrees[pair[1]] = append(outDegrees[pair[1]], pair[0])
	}

	q := []int{}
	// 把所有 in-degree 等於 0 的都裝進 queue 裡面
	for i, n := range numIndegrees {
		if n == 0 {
			q = append(q, i)
		}
	}

	// BFS
	for len(q) > 0 {
		size := len(q)
		for ; size > 0; size-- {
			i := q[0]
			q = q[1:]

			// 從所有 in-degree 為 0 的節點開始尋訪所有能走到的節點
			// 把這條 edge 斷掉後看是不是變成邊緣節點(in-degree=0)
			// 是的話就加進去繼續慢慢剝離, 直到所有邊緣節點都被拔完
			for _, next := range outDegrees[i] {
				numIndegrees[next]--
				if numIndegrees[next] == 0 {
					q = append(q, next)
				}
			}
		}
	}

	// 檢查閉環
	for _, n := range numIndegrees {
		// 剩下這些 in-degree 不為 0 的就是構成閉環的節點們了
		if n > 0 {
			return false
		}
	}
	return true
}
