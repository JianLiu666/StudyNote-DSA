package p00118

// Time Complexity: O(n^2)
// Space Complexity: O(1)
func getRow(rowIndex int) []int {
	row := make([]int, rowIndex+1)
	row[0] = 1

	for i := 1; i <= rowIndex; i++ {
		for j := i; j >= 1; j-- {
			row[j] += row[j-1]
		}
	}

	return row
}
