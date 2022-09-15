package p00052

// Time Complexity: O(n!)
// Space Complexity: O(n)
func totalNQueens(n int) (ans int) {
	columns := make([]bool, n)
	diagonals1 := make([]bool, n*2-1) // 紀錄左上到右下的對角線
	diagonals2 := make([]bool, n*2-1) // 紀錄右上到左下的對角線

	var dfs func(int)
	dfs = func(row int) {
		if row == n {
			ans++
			return
		}

		for col, hasQueen := range columns {
			d1 := row - col + n - 1
			d2 := row + col

			// pruning, 只要任何一個位置已經存在皇后就不用再繼續往下查找了
			if hasQueen || diagonals1[d1] || diagonals2[d2] {
				continue
			}

			columns[col] = true
			diagonals1[d1] = true
			diagonals2[d2] = true

			dfs(row + 1)

			columns[col] = false
			diagonals1[d1] = false
			diagonals2[d2] = false
		}
	}

	dfs(0)
	return
}
