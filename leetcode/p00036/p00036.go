package p00036

// Time Complexity: O(n)
// Space Complexity: O(1)
func isValidSudoku(board [][]byte) bool {
	rows := [9][9]int{}
	cols := [9][9]int{}
	recs := [9][9]int{}

	for row, columns := range board {
		for col, ch := range columns {
			if ch == '.' {
				continue
			}

			num := ch - '0' - 1
			rows[row][num]++
			if rows[row][num] > 1 {
				return false
			}

			cols[col][num]++
			if cols[col][num] > 1 {
				return false
			}

			rec := (row/3)*3 + (col / 3)
			recs[rec][num]++
			if recs[rec][num] > 1 {
				return false
			}
		}
	}

	return true
}
