package p00048

// Time Complexity: O(n), where n is the number of cells in matrix
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
	for i := 0; i < size; i++ {
		for j := i; j < size; j++ {
			matrix[j][i], matrix[i][j] = matrix[i][j], matrix[j][i]
		}
	}

	// reflect
	for c := 0; c < size/2; c++ {
		for r := 0; r < size; r++ {
			matrix[r][c], matrix[r][size-1-c] = matrix[r][size-1-c], matrix[r][c]
		}
	}
}
