package p00227

import (
	"regexp"
	"strconv"
)

// Time Complexity: O(n)
// Space Complexity: O(n)
func calculate(s string) int {
	result := 0

	// 移除所有空白符號
	re := regexp.MustCompile(`\s+`)
	s = re.ReplaceAllString(s, "")

	// trick, 因為題目說明只會有正數出現, 讓第一個數字的表達方式與後面格式一致
	s = "+" + s

	nums := []int{}

	for i := 0; i < len(s); i++ {
		j := i + 1
		for j < len(s) && s[j] >= '0' && s[j] <= '9' {
			j++
		}
		n, _ := strconv.Atoi(s[i+1 : j])
		switch s[i] {
		case '+':
			nums = append(nums, n)
		case '-':
			nums = append(nums, -n)
		case '*':
			nums[len(nums)-1] *= n
		case '/':
			nums[len(nums)-1] /= n
		}
		i = j - 1
	}

	for _, n := range nums {
		result += n
	}

	return result
}
