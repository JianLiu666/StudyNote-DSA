package p00221

func maximalSquare(matrix [][]byte) int {
	result := 0

	board := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		board[i] = make([]int, len(matrix[i]))
	}

	for row, columns := range matrix {
		for col, val := range columns {
			if row == 0 || col == 0 || val == '0' {
				board[row][col] = int(val - '0')
			} else {
				minimum := min(min(board[row-1][col], board[row][col-1]), board[row-1][col-1])
				board[row][col] = minimum + 1
			}
			result = max(result, board[row][col])
		}
	}

	return result * result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/**
問題分析
找到最大的 square, 所以每次只要找 matrix[row-1][col] matrix[row][col-1] 都大於 1 的時候

1 1 1 1    1 1 1 1
1 1 1 1 => 1 4 4 4
1 1 1 1    1 4 9 9
1 1 1 1

["1","1","1","1","0"],
["1","1","1","1","0"],
["1","1","1","1","1"],
["1","1","1","1","1"],
["0","0","1","1","1"]]

[
    ["1","1","1","1","0"]
    ["1","1","1","1","0"]
    ["1","1","1","1","1"]
    ["1","1","1","1","1"]
    ["0","0","1","1","1"]]
*/
