package p00315

// Time Complexity: O(nlogn)
// Space Complexity: O(n)
func countSmaller(nums []int) []int {
	if len(nums) == 1 {
		return []int{0}
	}

	counter := make([][]int, len(nums))
	for original_index, value := range nums {
		counter[original_index] = []int{value, original_index, 0}
	}

	counter = merge_and_sort(counter)
	res := make([]int, len(counter))
	for _, info := range counter {
		res[info[1]] = info[2]
	}

	return res
}

func merge_and_sort(nums [][]int) [][]int {
	if len(nums) == 1 {
		return nums
	}

	mid := len(nums) / 2
	l_arr := merge_and_sort(nums[:mid])
	r_arr := merge_and_sort(nums[mid:])

	res := make([][]int, len(nums))
	current := 0
	l_cur := 0
	r_cur := 0

	accumulator := 0
	for l_cur < len(l_arr) && r_cur < len(r_arr) {
		if l_arr[l_cur][0] > r_arr[r_cur][0] {
			accumulator++
			res[current] = r_arr[r_cur]
			r_cur++
		} else {
			// 先將目前為止的累加結果寫回 left element
			l_arr[l_cur][2] += accumulator
			res[current] = l_arr[l_cur]
			l_cur++
		}
		current++
	}

	for l_cur < len(l_arr) {
		// 將最後的累加結果寫回所有還在 left array 的 elements
		l_arr[l_cur][2] += accumulator
		res[current] = l_arr[l_cur]
		current++
		l_cur++
	}

	for r_cur < len(r_arr) {
		res[current] = r_arr[r_cur]
		current++
		r_cur++
	}

	return res
}
