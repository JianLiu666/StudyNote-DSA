package sorting

// Time Complexity: O(nlogn)
// Space Complexity: O(1)
func heapsort(nums []int) {
	// heapify part, build max-heap
	end := len(nums) - 1
	for i := end; i >= 0; i-- {
		shiftdown(nums, i, end)
	}

	// sorting part
	for i := 0; i < len(nums); i++ {
		nums[0], nums[end] = nums[end], nums[0]
		end--
		shiftdown(nums, 0, end)
	}
}

func shiftdown(nums []int, start, end int) {
	parent := start
	for {
		// get index of left child first
		child := parent*2 + 1
		if child > end {
			break
		}

		// right child exists and it greater than left child
		if child+1 <= end && nums[child] < nums[child+1] {
			child++
		}

		// if child is greater than parent, swap them
		if nums[parent] < nums[child] {
			nums[parent], nums[child] = nums[child], nums[parent]
			parent = child
		} else {
			break
		}
	}
}
