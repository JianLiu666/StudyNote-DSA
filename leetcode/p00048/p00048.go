package p00048

// Time Complexity: O(size), where n is the number of cells in matrix
// Space Complexity: O(1)
func rotate_bruteforce(matrix [][]int) {
	end := len(matrix) - 1

	for r := 0; r < end; r++ {
		for c := r; c < end-r; c++ {
			var tmp int
			matrix[c][end-r], tmp = matrix[r][c], matrix[c][end-r]
			matrix[end-r][end-c], tmp = tmp, matrix[end-r][end-c]
			matrix[end-c][r], tmp = tmp, matrix[end-c][r]
			matrix[r][c] = tmp
		}
	}
}

// Time Complexity: O(n), where n is the number of cells in matrix
// Space Complexity: O(1)
func rotate_linearalgebra(matrix [][]int) {
	size := len(matrix)

	// transpose
	// 從左上到右下斜線翻轉
	for row := 0; row < size; row++ {
		for col := 0; col < row; col++ {
			matrix[row][col], matrix[col][row] = matrix[col][row], matrix[row][col]
		}
	}

	// reflect
	for row := 0; row < size; row++ {
		for col := 0; col < size/2; col++ {
			matrix[row][col], matrix[row][size-1-col] = matrix[row][size-1-col], matrix[row][col]
		}
	}
}
