package p00179

import (
	"sort"
	"strconv"
)

// Time Complexity: O(nlogn)
// Space Complexity: O(n)
func largestNumber(nums []int) string {
	result := ""

	strs := []string{}
	for _, n := range nums {
		strs = append(strs, strconv.Itoa(n))
	}

	sort.Slice(strs, func(i, j int) bool {
		return strs[i]+strs[j] > strs[j]+strs[i]
	})

	for _, s := range strs {
		result += s
	}

	// 確保輸出格式符合題目規定, 即不會有 "00" 這種答案出現
	cursor := 0
	for cursor < len(result) && result[cursor] == '0' {
		cursor++
	}

	if cursor == len(result) {
		return "0"
	}

	return result[cursor:]
}
