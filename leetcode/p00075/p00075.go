package p00075

// Time Complexity: O(n)
// Space Complexity: O(1)
func sortColors(nums []int) {
	l, m, h := 0, 0, len(nums)-1

	// 目標是將所有數字排序完畢
	for m <= h {
		if nums[m] == 0 {
			nums[l], nums[m] = nums[m], nums[l]
			// 確定 0 已經被擺到最前面, 可以放心往後移
			l++
			m++
		} else if nums[m] == 1 {
			// 因此 l 會始終停留在第一個 1 的位置上, 等待遇到下一個 0 時跟 m 做 swap
			m++
		} else if nums[m] == 2 {
			// 把 2 擺到最後面, 但很有可能換回來的還是 2
			nums[m], nums[h] = nums[h], nums[m]
			h--
		}
	}
}
