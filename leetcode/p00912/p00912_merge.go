package p00912

// Time Complexity: O(nlogn)
// Space Complexity: O(n)
func sortArray_mergeSort_topDown(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}

	mid := len(nums) / 2

	arr1, arr2 := sortArray_mergeSort_topDown(nums[:mid]), sortArray_mergeSort_topDown(nums[mid:])
	start1, len1 := 0, len(arr1)
	start2, len2 := 0, len(arr2)

	res := make([]int, len1+len2)
	cur := 0

	for start1 < len1 && start2 < len2 {
		if arr1[start1] < arr2[start2] {
			res[cur] = arr1[start1]
			start1++
		} else {
			res[cur] = arr2[start2]
			start2++
		}
		cur++
	}
	for start1 < len1 {
		res[cur] = arr1[start1]
		start1++
		cur++
	}
	for start2 < len2 {
		res[cur] = arr2[start2]
		start2++
		cur++
	}

	return res
}

// Time Complexity: O(nlogn)
// Space Complexity: O(1)
func sortArray_mergeSort_bottomUp(nums []int) []int {
	size := len(nums)
	tmp := make([]int, size)

	for seg := 1; seg < size; seg += seg {
		for start := 0; start < size; start += seg * 2 {
			low, mid, high := start, min(start+seg, size), min(start+seg*2, size)
			cur := low
			start1, len1 := low, mid
			start2, len2 := mid, high
			for start1 < len1 && start2 < len2 {
				if nums[start1] < nums[start2] {
					tmp[cur] = nums[start1]
					start1++
				} else {
					tmp[cur] = nums[start2]
					start2++
				}
				cur++
			}
			for start1 < len1 {
				tmp[cur] = nums[start1]
				start1++
				cur++
			}
			for start2 < len2 {
				tmp[cur] = nums[start2]
				start2++
				cur++
			}
		}
		nums, tmp = tmp, nums
	}

	return nums
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
