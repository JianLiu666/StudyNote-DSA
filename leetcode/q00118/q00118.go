package q00118

// Time Complexity: O(n^2)
// Space Complexity: O(1)
func generate(numRows int) [][]int {
	levels := make([][]int, numRows)
	for row := 0; row < numRows; row++ {
		for col := 0; col <= row; col++ {
			levels[row] = append(levels[row], 1)
		}
	}

	if numRows <= 2 {
		return levels
	}

	for row := 2; row < numRows; row++ {
		for col := 1; col < row; col++ {
			levels[row][col] = levels[row-1][col-1] + levels[row-1][col]
		}
	}

	return levels
}
