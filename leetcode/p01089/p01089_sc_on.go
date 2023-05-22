package p01089

// Time Complexity: O(n)
// Space Complexity: O(n)
func duplicateZeros_sc_on(arr []int) {
	size := len(arr)
	copyright := make([]int, size)
	cursor := 0

	// 將 arr 的資料暫存一份在 copyright
	for i := 0; i < size && cursor < size; i++ {
		copyright[cursor] = arr[i]

		// 按照題目要求, 當遇到 0 的時候需要多複製一次
		if arr[i] == 0 && cursor < size-1 {
			cursor++
			copyright[cursor] = 0
		}
		cursor++
	}

	// 將 copyright 的結果對應回 arr
	for i := 0; i < size; i++ {
		arr[i] = copyright[i]
	}
}
