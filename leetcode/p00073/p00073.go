package p00073

// Time Complexity: O(mn), where m is the height of matrix, n is the width of matrix
// Space Complexity: O(mn)
func setZeroes(matrix [][]int) {
	rows := make([]bool, len(matrix))
	cols := make([]bool, len(matrix[0]))

	for row, columns := range matrix {
		for col, val := range columns {
			if val == 0 {
				rows[row] = true
				cols[col] = true
			}
		}
	}

	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[0]); col++ {
			if rows[row] || cols[col] {
				matrix[row][col] = 0
			}
		}
	}
}

// Time Complexity: O(mn), where m is the height of matrix, n is the width of matrix
// Space Complexity: O(1)
func setZeroes_optimize(matrix [][]int) {
	max_row, max_col := len(matrix), len(matrix[0])
	col0 := false

	// 遍歷整個 matrix
	for _, columns := range matrix {
		// 逐一檢查 col0 是否有 0
		if columns[0] == 0 {
			col0 = true
		}
		for col := 1; col < max_col; col++ {
			// 如果該 columns 內出現 0, 那就把 row0 對應位置的 row0[col] 標記成 0
			if columns[col] == 0 {
				columns[0] = 0
				matrix[0][col] = 0
			}
		}
	}

	// 反向遍歷整個 matrix
	// row[0] 是我們的參照物, 為了避免 row[0] 有 0 導致資訊錯誤, 所以從尾巴做回來
	for row := max_row - 1; row >= 0; row-- {
		for col := 1; col < max_col; col++ {
			// 如果相同 column 或 row 的頭為 0, 那該位置就要被標記成 0
			if matrix[row][0] == 0 || matrix[0][col] == 0 {
				matrix[row][col] = 0
			}
		}
		if col0 {
			matrix[row][0] = 0
		}
	}
}
