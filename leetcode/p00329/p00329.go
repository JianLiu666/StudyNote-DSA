package p00329

var Directions [4][2]int = [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func longestIncreasingPath(matrix [][]int) int {
	result := 0

	visited := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		visited[i] = make([]int, len(matrix[i]))
	}

	for row, columns := range matrix {
		for col := range columns {
			dfs(matrix, visited, row, col, &result)
		}
	}

	return result
}

func dfs(matrix, visited [][]int, row, col int, result *int) int {
	if visited[row][col] > 0 {
		return visited[row][col]
	}

	visited[row][col] = 1
	maxCount := 0
	for _, offset := range Directions {
		nextRow, nextCol := row+offset[0], col+offset[1]
		if nextRow < 0 || nextRow >= len(matrix) || nextCol < 0 || nextCol >= len(matrix[row]) {
			continue
		}
		if matrix[row][col] >= matrix[nextRow][nextCol] {
			continue
		}
		maxCount = max(maxCount, dfs(matrix, visited, nextRow, nextCol, result))
	}
	visited[row][col] += maxCount
	*result = max(*result, visited[row][col])

	return visited[row][col]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
