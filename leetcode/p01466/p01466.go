package p01466

func minReorder(n int, connections [][]int) int {
	result := 0

	visited := make([]bool, n)
	adlist := make([][]int, n)
	for _, pair := range connections {
		// 用正數代表 outgoing, 負數代表 incoming
		adlist[pair[0]] = append(adlist[pair[0]], pair[1])
		adlist[pair[1]] = append(adlist[pair[1]], -pair[0])
	}

	dfs(adlist, visited, 0, &result)

	return result
}

func dfs(adlist [][]int, visited []bool, idx int, result *int) {
	visited[idx] = true
	for _, next := range adlist[idx] {
		if next <= 0 {
			// incoming, 我們希望的
			next = abs(next)
			if !visited[next] {
				dfs(adlist, visited, next, result)
			}
		} else {
			// outgoing, 需要轉換的
			if !visited[next] {
				*result++
				dfs(adlist, visited, next, result)
			}
		}
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
