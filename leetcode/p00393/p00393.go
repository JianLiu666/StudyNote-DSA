package p00393

// Time Complexity: O(n)
// Space Complexity: O(1)
func validUtf8(data []int) bool {
	cur := 0

	for cur < len(data) {
		size := parseSequence(data[cur])
		if size == -1 || cur+size > len(data) {
			return false
		}
		if !validSequence(data[cur : cur+size]) {
			return false
		}
		cur += size
	}

	return true
}

func parseSequence(data int) int {
	data >>= 3
	if data == 0b11110 {
		return 4
	}
	data >>= 1
	if data == 0b1110 {
		return 3
	}
	data >>= 1
	if data == 0b110 {
		return 2
	}
	data >>= 2
	if data == 0 {
		return 1
	}

	return -1
}

func validSequence(data []int) bool {
	for i := 1; i < len(data); i++ {
		if data[i]>>6 != 0b10 {
			return false
		}
	}
	return true
}
