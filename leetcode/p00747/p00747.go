package p00747

// Time Complexity: O(n)
// Space Complexity: O(1)
func dominantIndex(nums []int) int {
	t1 := nums[0]
	t2 := nums[1]
	idx := 0
	if t2 > t1 {
		t1, t2 = t2, t1
		idx = 1
	}

	for i := 2; i < len(nums); i++ {
		if nums[i] > t1 {
			t1, t2 = nums[i], t1
			idx = i
		} else if nums[i] > t2 {
			t2 = nums[i]
		}
	}

	if t1 >= t2*2 {
		return idx
	}

	return -1
}
