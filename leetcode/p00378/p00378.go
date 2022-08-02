package p00378

// Time Complexity: O(nlogn)
// Space Complexity: O(1)
func kthSmallest(matrix [][]int, k int) int {
	maxRow, maxCol := len(matrix)-1, len(matrix[0])-1
	low, high := matrix[0][0], matrix[maxRow][maxCol]

	for low <= high {
		mid := low + (high-low)/2
		count := 0
		col := maxCol
		for row := 0; row <= maxRow; row++ {
			for col >= 0 && matrix[row][col] > mid {
				col--
			}
			count += col + 1
		}

		if count < k {
			low = mid + 1
		} else {
			high = mid
		}
	}

	return low
}
