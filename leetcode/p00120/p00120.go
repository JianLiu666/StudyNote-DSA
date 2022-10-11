package p00120

// Time Complexity: O(n^2)
// Space Complexity: O(1)
func minimumTotal(triangle [][]int) int {
	for r := 1; r < len(triangle); r++ {
		for c := 0; c <= r; c++ {
			if c == 0 {
				// 左邊界處理
				triangle[r][c] += triangle[r-1][0]
			} else if c == r {
				// 右邊界處理
				triangle[r][c] += triangle[r-1][c-1]
			} else {
				// 一般處理流程, 找到本次加總後的最小值
				triangle[r][c] += min(triangle[r-1][c-1], triangle[r-1][c])
			}
		}
	}

	row := len(triangle)
	// 找出最後一層的最小值
	res := triangle[row-1][0]
	for i := 1; i < row; i++ {
		if res > triangle[row-1][i] {
			res = triangle[row-1][i]
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
