package p00415

import "strings"

// Time Complexity: O(n), where n is the bigger length of number
// Space Complexity: O(n)
func addStrings(num1 string, num2 string) string {
	c1, c2 := len(num1)-1, len(num2)-1

	var str strings.Builder
	var carried byte

	for c1 >= 0 || c2 >= 0 || carried > 0 {
		var num byte = '0'
		if c1 >= 0 {
			num += num1[c1] - '0'
			c1--
		}
		if c2 >= 0 {
			num += num2[c2] - '0'
			c2--
		}
		num += carried
		carried = 0
		if num > '9' {
			carried = 1
			num -= 10
		}
		str.WriteByte(num)
	}

	reversed := str.String()
	str.Reset()
	for i := len(reversed) - 1; i >= 0; i-- {
		str.WriteByte(reversed[i])
	}

	return str.String()
}
