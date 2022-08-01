package p00542

// Time Complexity: O(mn), where m is length of rows, and n is length of columns
// Space Complexity: O(1)
func updateMatrix_dp(mat [][]int) [][]int {
	nRows, nCols := len(mat), len(mat[0])
	maxDistance := nRows + nCols

	for r := 0; r < nRows; r++ {
		for c := 0; c < nCols; c++ {
			if mat[r][c] == 0 {
				continue
			}

			left, top := maxDistance, maxDistance
			if r-1 >= 0 {
				top = mat[r-1][c]
			}
			if c-1 >= 0 {
				left = mat[r][c-1]
			}
			mat[r][c] = min(top, left) + 1
		}
	}

	for r := nRows - 1; r >= 0; r-- {
		for c := nCols - 1; c >= 0; c-- {
			if mat[r][c] == 0 {
				continue
			}

			right, bottom := maxDistance, maxDistance
			if r+1 < nRows {
				bottom = mat[r+1][c]
			}
			if c+1 < nCols {
				right = mat[r][c+1]
			}
			mat[r][c] = min(mat[r][c], min(bottom, right)+1)
		}
	}

	return mat
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
