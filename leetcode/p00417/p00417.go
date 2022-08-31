package p00417

var MAX_ROW, MAX_COL int
var pos [][]int = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

// Time Complexity: O((mn)^2), where m is the height of heights, n is the width of heights
// Space Complexity: O(mn)
func pacificAtlantic(heights [][]int) [][]int {
	MAX_ROW, MAX_COL = len(heights), len(heights[0])

	pacific := make([][]bool, MAX_ROW)
	atlantic := make([][]bool, MAX_ROW)
	for i := 0; i < MAX_ROW; i++ {
		pacific[i] = make([]bool, MAX_COL)
		atlantic[i] = make([]bool, MAX_COL)
	}

	// 沿著左右邊界追蹤水的源頭
	for row := 0; row < MAX_ROW; row++ {
		dfs(heights, pacific, 0, row, 0)
		dfs(heights, atlantic, 0, row, MAX_COL-1)
	}

	// 沿著上下邊界追蹤水的源頭
	for col := 0; col < MAX_COL; col++ {
		dfs(heights, pacific, 0, 0, col)
		dfs(heights, atlantic, 0, MAX_ROW-1, col)
	}

	res := [][]int{}
	for row := 0; row < MAX_ROW; row++ {
		for col := 0; col < MAX_COL; col++ {
			if pacific[row][col] && atlantic[row][col] {
				res = append(res, []int{row, col})
			}
		}
	}
	return res
}

func dfs(heights [][]int, visited [][]bool, height, row, col int) {
	// 邊界判斷
	if row < 0 || row >= MAX_ROW || col < 0 || col >= MAX_COL {
		return
	}
	// 遇到已經尋訪過的路提早剪枝
	if visited[row][col] {
		return
	}
	// 確保海拔高度比前一個位置還要高, 才能使雨水往低處流
	if heights[row][col] < height {
		return
	}

	visited[row][col] = true

	// 往四周繼續探尋是否還有更高海拔的源頭
	for _, p := range pos {
		dfs(heights, visited, heights[row][col], row+p[0], col+p[1])
	}
}
