package p01657

import "sort"

// Time Complexity: O(nlogn), where n is the length of word
// Space Complexity: O(1)
func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	memo1 := make([]int, 26)
	memo2 := make([]int, 26)
	for i := 0; i < len(word1); i++ {
		memo1[word1[i]-'a']++
		memo2[word2[i]-'a']++
	}

	// 先比較有的字母都一樣
	for i := 0; i < len(memo1); i++ {
		if (memo1[i] > 0 && memo2[i] == 0) || (memo1[i] == 0 && memo2[i] > 0) {
			return false
		}
	}

	// 在比較出現數量也都一樣
	sort.Ints(memo1)
	sort.Ints(memo2)
	for i := 0; i < len(memo1); i++ {
		if memo1[i] != memo2[i] {
			return false
		}
	}

	return true
}

/**
分析題目:
只要每個字母的出現頻率一樣, 就可以用 operation1 換到一樣的位置 -> i.e. 位置不重要
只要統計每個字母的出現頻率後, 重新排序一遍, 如果一模一樣就是答案

example3.

231
123
*/
