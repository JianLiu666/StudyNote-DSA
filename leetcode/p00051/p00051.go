package p00051

import "strings"

func solveNQueens(n int) [][]string {
	result := [][]string{}

	// 初始化盤面
	board := make([][]string, n)
	for row := 0; row < n; row++ {
		board[row] = make([]string, n)
		for col := 0; col < n; col++ {
			board[row][col] = "."
		}
	}

	// memorization, 記錄皇后放過的位置
	rowVisited := map[int]bool{}
	colVisited := map[int]bool{}
	leftDiagonalVisited := map[int]bool{}
	rightDiagonalVisited := map[int]bool{}

	dfs(board, 0, 0, n, 0, rowVisited, colVisited, leftDiagonalVisited, rightDiagonalVisited, &result)

	return result
}

func dfs(board [][]string, row, col, maxQueens, nQueens int, rowVisited, colVisited, leftDiagonalVisited, rightDiagonalVisited map[int]bool, result *[][]string) {
	if nQueens == maxQueens {
		// early pruning
		solution := []string{}
		for _, column := range board {
			solution = append(solution, strings.Join(column, ""))
		}
		*result = append(*result, solution)
		return
	}
	if row == maxQueens {
		// 已經走到超出棋盤大小就結束
		// 同時也代表這個下法無解
		return
	}

	nextRow, nextCol := next(maxQueens, row, col)

	// 記錄當前狀態, 退回時要用到
	rowPrevStatus := rowVisited[row]
	colPrevStatus := colVisited[col]
	leftDiagonalPrevStatus := leftDiagonalVisited[row-col]
	rightDiagonalPrevStatus := rightDiagonalVisited[row+col]

	// 檢查所有方向確認能不能放下皇后
	if !rowPrevStatus && !colPrevStatus && !leftDiagonalPrevStatus && !rightDiagonalPrevStatus {
		board[row][col] = "Q"

		rowVisited[row] = true
		colVisited[col] = true
		leftDiagonalVisited[row-col] = true
		rightDiagonalVisited[row+col] = true

		dfs(board, nextRow, nextCol, maxQueens, nQueens+1, rowVisited, colVisited, leftDiagonalVisited, rightDiagonalVisited, result)

		rowVisited[row] = rowPrevStatus
		colVisited[col] = colPrevStatus
		leftDiagonalVisited[row-col] = leftDiagonalPrevStatus
		rightDiagonalVisited[row+col] = rightDiagonalPrevStatus
	}

	board[row][col] = "."
	dfs(board, nextRow, nextCol, maxQueens, nQueens, rowVisited, colVisited, leftDiagonalVisited, rightDiagonalVisited, result)
}

func next(n, row, col int) (int, int) {
	if col+1 >= n {
		return row + 1, 0
	}
	return row, col + 1
}
