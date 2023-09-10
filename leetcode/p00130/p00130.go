package p00130

var directions [4][2]int = [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func solve(board [][]byte) {
	for row, column := range board {
		for col, val := range column {
			if (row == 0 || col == 0 || row == len(board)-1 || col == len(board[row])-1) && val == 'O' {
				dfs(board, row, col)
			}
		}
	}

	for row, column := range board {
		for col, val := range column {
			if val == 'V' {
				board[row][col] = 'O'
			} else {
				board[row][col] = 'X'
			}
		}
	}
}

func dfs(board [][]byte, row, col int) {
	board[row][col] = 'V'

	for _, offset := range directions {
		nr, nc := row+offset[0], col+offset[1]
		if nr < 0 || nr == len(board) || nc < 0 || nc == len(board[row]) || board[nr][nc] != 'O' {
			continue
		}
		dfs(board, nr, nc)
	}
}

/**
解題想法:
第一次從所有邊緣是 'O' 的位置開始往內走訪並標記成 'V', 第二次遍歷 board 時把所有被標記成 'V' 的轉換成 'O', 其餘轉換成 'X'
*/
