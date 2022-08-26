package p00869

import "strings"

// Time Complexity: O(n), where n is the number of digits
// Space Complexity: O(1)
func reorderedPowerOf2(n int) bool {
	str := num2str(n)
	num := 1
	for i := 1; i <= 30; i++ {
		if str == num2str(num) {
			return true
		}
		num *= 2
	}

	return false
}

func num2str(num int) string {
	// 表示 0-9
	counter := [10]int{}
	for num > 0 {
		var digit int
		digit, num = num%10, num/10
		counter[digit]++
	}

	var str strings.Builder
	for _, digit := range counter {
		str.WriteByte(byte(digit + 48))
	}

	return str.String()
}
