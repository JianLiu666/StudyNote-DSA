package p01680

// Time Complexity: O(n)
// Space Complexity: O(1)
func concatenatedBinary(n int) int {
	sum := 0
	shift := 0
	mod := 1000000007

	for i := 1; i <= n; i++ {
		// 判斷是否是2的次方, 用來決定要 right shift 的次數
		if i&(i-1) == 0 {
			shift++
		}
		sum <<= shift
		sum += i
		sum %= mod
	}

	return sum
}
