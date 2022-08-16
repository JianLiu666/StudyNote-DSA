package p00074

// Time Complexity: O(m+n), where m is the length of matrix, n is the length of matrix[0]
// Space Complexity: O(1)
func searchMatrix(matrix [][]int, target int) bool {
	r, c := len(matrix)-1, 0

	for r >= 0 && c < len(matrix[0]) {
		if matrix[r][c] == target {
			return true
		}
		if matrix[r][c] > target {
			r--
		} else {
			c++
		}
	}

	return false
}
