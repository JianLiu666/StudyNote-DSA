package p01466

func minReorder_bruteforce(n int, connections [][]int) int {
	result := 0

	// step.1 建立無向圖
	undirected := map[int]map[int]bool{}
	for _, pair := range connections {
		if _, exists := undirected[pair[0]]; !exists {
			undirected[pair[0]] = map[int]bool{}
		}
		if _, exists := undirected[pair[1]]; !exists {
			undirected[pair[1]] = map[int]bool{}
		}
		undirected[pair[0]][pair[1]] = true
		undirected[pair[1]][pair[0]] = true
	}

	// step.2 建立有向圖
	dag := map[int]map[int]bool{}
	q := []int{0}

	for len(q) > 0 {
		size := len(q)
		for ; size > 0; size-- {
			idx := q[0]
			q = q[1:]

			if _, exists := dag[idx]; exists {
				continue
			}
			dag[idx] = map[int]bool{}

			for next := range undirected[idx] {
				if _, exists := dag[next]; exists {
					continue
				}
				dag[idx][next] = true
				q = append(q, next)
			}
		}
	}

	// step.3 跟 connections 比較差異
	for _, pair := range connections {
		if dag[pair[0]][pair[1]] {
			result++
		}
	}

	return result
}

/**
step.1 建構一個無向的 graph
step.2 BFS 走訪過一次建構一個有向的 graph
step.3 用有向的 graph 跟 connections 比較一次計算差異

n = 6, connections = [[0,1],[1,3],[2,3],[4,0],[4,5]]
*/
