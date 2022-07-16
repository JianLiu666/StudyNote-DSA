package p00498

// Time Complexity: O(n)
// Space Complexity: O(n)
func findDiagonalOrder(mat [][]int) []int {
	m := 0
	mx := len(mat)
	n := 0
	nx := len(mat[0])

	// 0: up-right, 1: down-left
	direction := 0

	res := make([]int, mx*nx)
	count := 0
	for m < mx && n < nx {
		res[count] = mat[m][n]
		count++

		if direction == 0 {
			if m-1 >= 0 && n+1 < nx {
				m--
				n++
			} else {
				direction = 1
				if n+1 < nx {
					n++
				} else {
					m++
				}
			}
		} else {
			if m+1 < mx && n-1 >= 0 {
				m++
				n--
			} else {
				direction = 0
				if m+1 < mx {
					m++
				} else {
					n++
				}
			}
		}
	}

	return res
}
