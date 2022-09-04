package p01582

// Time Complexity: O(mn), where m is the height of mat, n is the width of mat
// Space Complexity: O(m+n)
func numSpecial(mat [][]int) int {
	rows := make([]int, len(mat))
	cols := make([]int, len(mat[0]))

	for row, columns := range mat {
		for col, val := range columns {
			if val == 1 {
				rows[row]++
				cols[col]++
			}
		}
	}

	count := 0
	for row, columns := range mat {
		for col, val := range columns {
			if val == 1 && rows[row] == 1 && cols[col] == 1 {
				count++
			}
		}
	}
	return count
}

// Time Complexity: O(mn), where m is the height of mat, n is the width of mat
// Space Complexity: O(1)
func numSpecial_min_sc(mat [][]int) int {
	for row, columns := range mat {
		cnt := 0
		// 因為 mat elements 只包含 0,1, 所以直接累加即可
		for _, val := range columns {
			cnt += val
		}
		// 避免 mat[0] 出現 1 時, 把自己也重複計算一遍
		if row == 0 {
			cnt--
		}
		if cnt > 0 {
			// 把出現 1 的位置更新到對應的 mat[0] 位置上
			for col, val := range columns {
				if val == 1 {
					mat[0][col] += cnt
				}
			}
		}
	}

	res := 0
	// 所有累計都已經在 mat[0] 上, 只要 mat[i][j] == 1 表示為符合題意的 special position
	for _, val := range mat[0] {
		if val == 1 {
			res++
		}
	}

	return res
}
