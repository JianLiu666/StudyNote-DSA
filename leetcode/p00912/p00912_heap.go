package p00912

// Time Complexity: O(nlogn)
// Space Complexity: O(1)
func sortArray_heapSort(nums []int) []int {
	size := len(nums)
	end := size - 1
	for i := end; i >= 0; i-- {
		shiftDown(nums, i, end)
	}

	for i := 0; i < size; i++ {
		nums[0], nums[end] = nums[end], nums[0]
		end--
		shiftDown(nums, 0, end)
	}

	return nums
}

func shiftDown(nums []int, start, end int) {
	parent := start
	for {
		child := parent*2 + 1
		if child > end {
			break
		}

		if child+1 <= end && nums[child+1] > nums[child] {
			child++
		}

		if nums[child] > nums[parent] {
			nums[child], nums[parent] = nums[parent], nums[child]
			parent = child
		} else {
			break
		}
	}
}
