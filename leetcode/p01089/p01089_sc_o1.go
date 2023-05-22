package p01089

// Time Complexity: O(n)
// Space Complexity: O(1)
func duplicateZeros_sc_o1(arr []int) {
	head := len(arr) - 1
	tail := len(arr) - 1
	cursor := 0

	// 從後往前, 根據 0 的數量定位 head, tail 的位置
	for cursor < head {
		if arr[cursor] == 0 {
			head--
		}
		cursor++

		// 如果 cursor 與 head 相遇時剛好是 0 的時候, 需要先在 tail 的位置補上 0
		// 並且把 head, tail 都往前移動一格, 否則最後當 tail 追到 head 時會多處理到一個 0
		if arr[head] == 0 && cursor == head {
			arr[tail] = 0
			head--
			tail--
		}
	}

	// 從後往前移動, 不斷把 head 指到的 value 複製到 tail 的位置
	// 直到沒有 0 或是移動到 array 的最前面
	for head > -1 && head != tail {
		arr[tail] = arr[head]
		if arr[head] == 0 {
			tail--
			arr[tail] = arr[head]
		}
		head--
		tail--
	}
}
