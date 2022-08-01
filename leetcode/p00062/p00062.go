package p00062

// Time Complexity: O(mn)
// Space Complexity: O(mn)
func uniquePaths(m int, n int) int {
	matrix := make([][]int, m)
	for row := 0; row < len(matrix); row++ {
		matrix[row] = make([]int, n)
	}

	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			sum := 1
			if row > 0 && col > 0 {
				sum = matrix[row-1][col] + matrix[row][col-1]
			}
			matrix[row][col] = sum
		}
	}

	return matrix[m-1][n-1]
}
