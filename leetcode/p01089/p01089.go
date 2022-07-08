package p01089

// Time Complexity: O(n)
// Space Complexity: O(1)
func duplicateZeros(arr []int) {
	head := len(arr) - 1
	tail := len(arr) - 1
	current := 0

	for current < head {
		if arr[current] == 0 {
			head--
		}
		current++

		if arr[head] == 0 && current == head {
			arr[tail] = 0
			head--
			tail--
		}
	}

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
