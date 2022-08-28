package p00059

// Time Complexity: O(n^2)
// Space Complexity: O(1)
func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	start, end := 1, n*n

	min_row, min_col := 1, 0
	max_row, max_col := n, n
	cur_row, cur_col := 0, 0

	// 1: right, 2: down, 3: left, 4: up
	direction := 1

	for start <= end {
		matrix[cur_row][cur_col] = start
		start++

		if direction == 1 {
			if cur_col+1 < max_col {
				// right
				cur_col++
			} else {
				// down
				direction = 2
				cur_row++
				max_col--
			}
		} else if direction == 2 {
			if cur_row+1 < max_row {
				// down
				cur_row++
			} else {
				// left
				direction = 3
				cur_col--
				max_row--
			}
		} else if direction == 3 {
			if cur_col-1 >= min_col {
				// left
				cur_col--
			} else {
				// up
				direction = 4
				cur_row--
				min_col++
			}
		} else {
			if cur_row-1 >= min_row {
				// up
				cur_row--
			} else {
				// right
				direction = 1
				cur_col++
				min_row++
			}
		}
	}

	return matrix
}
