package p00287

// Time Complexity: O(nlogn)
// Space Complexity: O(1)
func findDuplicate_binarysearch(nums []int) int {
	low, high := 1, len(nums)-1

	for low < high {
		mid := low + (high-low)/2
		count := 0
		for _, num := range nums {
			if num <= mid {
				count++
			}
		}

		if count > mid {
			high = mid
		} else {
			low = mid + 1
		}
	}

	return low
}

// Time Complexity: O(n)
// Space Complexity: O(1)
func findDuplicate_cycledetection(nums []int) int {
	slow := nums[0]
	fast := nums[nums[0]]

	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}

	anchor := 0
	for slow != anchor {
		slow = nums[slow]
		anchor = nums[anchor]
	}

	return anchor
}
