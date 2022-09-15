package p00037

func solveSudoku(board [][]byte) {
	// 為了方便後續計算, index 從 1 開始記錄
	rows := make([]int, 10)
	cols := make([]int, 10)
	recs := make([]int, 10)

	for i := 1; i <= 9; i++ {
		rows[i] = 0b1111111110
		cols[i] = 0b1111111110
		recs[i] = 0b1111111110
	}

	// 初始化可用數字
	for row, columns := range board {
		for col, val := range columns {
			if val == '.' {
				continue
			}

			xor := 1 << (val - '0')
			rows[col+1] ^= xor
			cols[row+1] ^= xor
			recs[(row/3)*3+col/3+1] ^= xor
		}
	}

	dfs(board, rows, cols, recs, 0, 0)
}

func dfs(board [][]byte, rows, cols, recs []int, r, c int) bool {
	if r == 9 {
		return true
	}

	nr, nc := next(r, c)

	// 如果是原本就存在的數字就直接跳過
	if board[r][c] != '.' {
		return dfs(board, rows, cols, recs, nr, nc)
	}

	// 嘗試所有可能
	for i := 1; i <= 9; i++ {
		xor := 1 << i
		// 只要這個數字填入後違反數獨規則就不用再繼續往下了
		if !valid(xor, r, c, rows, cols, recs) {
			continue
		}

		board[r][c] = byte('0' + i)
		rec := (r/3)*3 + c/3

		rows[c+1] ^= xor
		cols[r+1] ^= xor
		recs[rec+1] ^= xor

		// 只要找到一組解就直接結束
		if dfs(board, rows, cols, recs, nr, nc) {
			return true
		}

		board[r][c] = '.'
		rows[c+1] ^= xor
		cols[r+1] ^= xor
		recs[rec+1] ^= xor
	}

	return false
}

func next(r, c int) (int, int) {
	if c < 8 {
		return r, c + 1
	}
	return r + 1, 0
}

func valid(num, r, c int, rows, cols, recs []int) bool {
	if rows[c+1]&num == 0 {
		return false
	}
	if cols[r+1]&num == 0 {
		return false
	}
	if recs[(r/3)*3+c/3+1]&num == 0 {
		return false
	}
	return true
}
