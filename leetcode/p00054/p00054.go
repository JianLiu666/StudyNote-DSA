package p00054

// Time Complexity: O(n)
// Space Complexity: O(n)
func spiralOrder(matrix [][]int) []int {
	m_min, m_cur, m_max := -1, 0, len(matrix)
	n_min, n_cur, n_max := -1, 0, len(matrix[0])

	res := make([]int, m_max*n_max)
	count := 0

	// 0: right, 1: down, 2: left, 3: up
	direction := 0
	for count < len(res) {
		res[count] = matrix[m_cur][n_cur]
		count++

		if direction == 0 {
			if n_cur+1 < n_max {
				n_cur++
			} else {
				direction = 1
				m_cur++
				m_min++
			}
		} else if direction == 1 {
			if m_cur+1 < m_max {
				m_cur++
			} else {
				direction = 2
				n_cur--
				n_max--
			}
		} else if direction == 2 {
			if n_cur-1 > n_min {
				n_cur--
			} else {
				direction = 3
				m_cur--
				m_max--
			}
		} else {
			if m_cur-1 > m_min {
				m_cur--
			} else {
				direction = 0
				n_cur++
				n_min++
			}
		}
	}

	return res
}
