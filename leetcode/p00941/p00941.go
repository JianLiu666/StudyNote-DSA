package p00941

// Time Complexity: O(n)
// Space Complexity: O(1)
func validMountainArray(arr []int) bool {
	head := 0
	tail := len(arr) - 1

	for head != tail {
		if arr[head] >= arr[head+1] && arr[tail-1] <= arr[tail] {
			return false
		}

		if arr[head+1] > arr[head] {
			head++
		}

		if arr[tail-1] > arr[tail] {
			tail--
		}
	}

	if head == 0 || tail == len(arr)-1 {
		return false
	}

	return true
}
