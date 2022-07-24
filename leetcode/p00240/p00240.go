package p00240

// Time Complexity: O(m+n)
// Space Complexity: O(1)
func searchMatrix(matrix [][]int, target int) bool {
	col := len(matrix) - 1
	row := 0

	for col >= 0 && col <= len(matrix)-1 && row >= 0 && row <= len(matrix[0])-1 {
		if matrix[col][row] == target {
			return true
		} else if matrix[col][row] < target {
			row++
		} else {
			col--
		}
	}

	return false
}

// Time Complexity: O(mlogn)
// Space Complexity: O(1)
func searchMatrix_binarysearch(matrix [][]int, target int) bool {
	for _, arr := range matrix {
		if target < arr[0] || target > arr[len(arr)-1] {
			continue
		}

		head := 0
		tail := len(matrix[0]) - 1
		for head <= tail {
			mid := head + (tail-head)/2
			if arr[mid] == target {
				return true
			} else if arr[mid] > target {
				tail = mid - 1
			} else {
				head = mid + 1
			}
		}
	}

	return false
}
