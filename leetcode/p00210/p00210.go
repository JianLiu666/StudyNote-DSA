package p00210

// Time Complexity: O(n+m), where n is the numCourses, m is the number of prerequisites
// Space Complexity: O(n+m), where n is the numCourses, m is the number of prerequisites
func findOrder(numCourses int, prerequisites [][]int) []int {
	graph := map[int][]int{}
	validCourses := map[int]bool{}

	// 初始化 graph
	for _, course := range prerequisites {
		if graph[course[0]] == nil {
			graph[course[0]] = make([]int, 0)
		}
		graph[course[0]] = append(graph[course[0]], course[1])
	}

	res := []int{}

	// 逐一檢查每個課程是否存在閉環迴路
	for i := 0; i < numCourses; i++ {
		if detectCycle(graph, map[int]bool{}, validCourses, i, &res) {
			return []int{}
		}
	}

	return res
}

func detectCycle(graph map[int][]int, visitedCourses, validCourses map[int]bool, node int, res *[]int) bool {
	// 已經確定合法的課程就不用在往下檢查了
	if validCourses[node] {
		return false
	}
	// 遇到本次尋訪曾經拜訪過的課程表示存在閉環迴路
	if visitedCourses[node] {
		return true
	}

	// preorder concept, 紀錄本次尋訪走過的課程
	visitedCourses[node] = true

	// DFS concept
	size := len(graph[node])
	for i := 0; i < size; i++ {
		// 如果已經發生閉環就可以提早結束了
		if detectCycle(graph, visitedCourses, validCourses, graph[node][i], res) {
			return true
		}
	}

	// 移除本次尋訪走過的課程, 避免影響到其他路徑尋訪
	visitedCourses[node] = false

	// 確定這個課程已經沒問題了
	validCourses[node] = true

	// 確定沒問題的課程就直接記錄下來
	*res = append(*res, node)

	return false
}
