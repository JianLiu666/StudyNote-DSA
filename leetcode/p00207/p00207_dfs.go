package p00207

// Time Complexity: O(n+m), where n is the numCourses, m is the number of prerequisites
// Space Complexity: O(n+m), where n is the numCourses, m is the number of prerequisites
func canFinish_dfs(numCourses int, prerequisites [][]int) bool {
	// record node status: 0: not visited yet, 1: valid, 2: unvalid
	visited := map[int]int{}

	indegrees := make([][]int, numCourses)
	for _, pair := range prerequisites {
		indegrees[pair[0]] = append(indegrees[pair[0]], pair[1])
	}

	// start from course 0 to latest
	for i := 0; i < numCourses; i++ {
		if !dfs(visited, indegrees, i) {
			return false
		}
	}

	return true
}

// @param visited 用來記錄節點狀態: 0:未訪問, 1:能結束, 2:不能結束/不確定
// @param indegrees 記錄所有節點的 in-degree
// @param 目前要處理的節點編號
//
// @return bool 是否出現閉環
func dfs(visited map[int]int, indegrees [][]int, i int) bool {
	if visited[i] == 1 {
		return true
	}
	if visited[i] == 2 {
		return false
	}
	// 還不確定到底能不能走訪完, 正要開始訪問先記成2
	// 一旦在這個節點訪問完畢之前又被訪問到, 那就代表一定是閉環
	visited[i] = 2
	for _, next := range indegrees[i] {
		// 只要提早發現閉環那就不用再繼續下去了 (early pruning)
		if !dfs(visited, indegrees, next) {
			return false
		}
	}
	// 從這個節點出發到結束之間沒有出現閉環
	visited[i] = 1
	return true
}
