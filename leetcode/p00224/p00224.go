package p00224

import "strconv"

// Time Complexity: O(n)
// Space Complexity: O(n)
func calculate(s string) int {
	result := 0

	s = preprocess(s)

	nums := []int{}
	signs := []int{}
	// 用來處理正負號, 因此可以視為每一次的操作都是 result 加上一個數字
	sign := 1

	for i := 0; i < len(s); i++ {
		if s[i] == '+' {
			sign = 1
		} else if s[i] == '-' {
			sign = -1
		} else if s[i] >= '0' && s[i] <= '9' {
			j := i + 1
			for j < len(s) && s[j] >= '0' && s[j] <= '9' {
				j++
			}
			n, _ := strconv.Atoi(s[i:j])
			result += sign * n
			i = j - 1
		} else if s[i] == '(' {
			// 遇到左括號表示需要將之前的結果先暫存起來, 讓計算從新累計
			nums = append(nums, result)
			signs = append(signs, sign)
			result = 0
		} else if s[i] == ')' {
			// 遇到右括號表示可以將現在的計算結果加回之前的暫存結果
			result *= signs[len(signs)-1]
			signs = signs[:len(signs)-1]
			result += nums[len(nums)-1]
			nums = nums[:len(nums)-1]
		}
	}

	return result
}

// 確保運算式內的每個數字之前都有符號表示正負
func preprocess(s string) string {
	if s[0] != '-' {
		s = "+" + s
	}

	tmp := ""
	for _, ch := range s {
		if ch == ' ' {
			continue
		}
		tmp += string(ch)
		if ch == '(' {
			tmp += "+"
		}
	}

	return tmp
}
