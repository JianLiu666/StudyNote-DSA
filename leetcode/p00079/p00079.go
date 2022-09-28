package p00079

// Time Complexity: O(mn*3^l), where m is the width of board, n is the height of board, l is the length of word
// Space Complexity: O(mn)
func exist(board [][]byte, word string) bool {
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	visited := [6][6]bool{}

	// tc: O(n)
	for row, columns := range board {
		// tc: O(m)
		for col, val := range columns {
			if val != word[0] {
				continue
			}
			if dfs(visited, board, row, col, directions, word, 0) {
				return true
			}
		}
	}

	return false
}

func dfs(visited [6][6]bool, board [][]byte, row, col int, directions [][]int, word string, idx int) bool {
	// 只要任何一個字母不吻合就提早中斷
	if board[row][col] != word[idx] {
		return false
	}
	if idx+1 == len(word) {
		return true
	}

	visited[row][col] = true
	for _, pos := range directions {
		y := row + pos[0]
		x := col + pos[1]
		// tc: O(3^l), 原因在於走過的路不能重複, 所以每次最多只會有三個方向可以前進
		if x < 0 || y < 0 || x >= len(board[0]) || y >= len(board) || visited[y][x] {
			continue
		}
		// 一旦找到完全相符的字串, 就可以提早終止剩餘操作
		res := dfs(visited, board, y, x, directions, word, idx+1)
		if res {
			return true
		}
	}
	visited[row][col] = false

	return false
}
