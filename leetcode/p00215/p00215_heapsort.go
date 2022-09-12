package p00215

// Time Complexity: O(nlogn), where n is the number of nums
// Space Complexity: O(1)
func findKthLargest_heapsort(nums []int, k int) int {
	end := len(nums) - 1

	// heapify part, build max-heap
	for i := end; i >= 0; i-- {
		shiftdown(nums, i, end)
	}

	res := nums[end]

	// sorting part
	for i := 1; i <= k; i++ {
		res = nums[0]
		nums[0], nums[end] = nums[end], nums[0]
		end--
		shiftdown(nums, 0, end)
	}

	return res
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
		if child+1 <= end && nums[child+1] > nums[child] {
			child++
		}

		// if child is greater than parent, swap them
		if nums[child] > nums[parent] {
			nums[child], nums[parent] = nums[parent], nums[child]
			parent = child
		} else {
			break
		}
	}
}
