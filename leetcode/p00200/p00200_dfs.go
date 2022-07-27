package p00200

// Time Complexity: O(m+n)
// Space Complexity: O(m)
func numIslands_dfs(grid [][]byte) int {
	count := 0
	for row, columns := range grid {
		for col, val := range columns {
			if val == '1' {
				count++
				dfs(grid, row, col)
			}
		}
	}

	return count
}

func dfs(grid [][]byte, row, col int) {
	grid[row][col] = '2'

	// up
	if row-1 >= 0 && grid[row-1][col] == '1' {
		dfs(grid, row-1, col)
	}

	// right
	if col+1 <= len(grid[0])-1 && grid[row][col+1] == '1' {
		dfs(grid, row, col+1)
	}

	// down
	if row+1 <= len(grid)-1 && grid[row+1][col] == '1' {
		dfs(grid, row+1, col)
	}

	// left
	if col-1 >= 0 && grid[row][col-1] == '1' {
		dfs(grid, row, col-1)
	}
}
