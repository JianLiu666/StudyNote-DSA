package p00336

// Time Complexity: O(n * m^2), where n is the length of words, m is the length of each word
// Space Complexity: O(n)
func palindromePairs(words []string) [][]int {
	result := [][]int{}

	// 紀錄 word 的反轉內容, 為了檢查回文
	memo := map[string]int{}
	for i, word := range words {
		memo[reverse(word)] = i
	}

	// 先用一個去重複的 map 過濾掉重複的可能
	set := map[[2]int]bool{}

	for i, word := range words {
		// 對每一個 word 進行處理, 找到其他 word 可以構成回文的部分, 也檢查本身是否屬於回文
		for divide := 0; divide <= len(word); divide++ {
			if j, exists := memo[word[:divide]]; exists && i != j && isPalindrome(word[divide:]) {
				set[[2]int{i, j}] = true
			}
			if j, exists := memo[word[divide:]]; exists && i != j && isPalindrome(word[:divide]) {
				set[[2]int{j, i}] = true
			}
		}
	}

	for pair := range set {
		result = append(result, []int{pair[0], pair[1]})
	}

	return result
}

func isPalindrome(str string) bool {
	for head, tail := 0, len(str)-1; head <= tail; head, tail = head+1, tail-1 {
		if str[head] != str[tail] {
			return false
		}
	}

	return true
}

func reverse(str string) string {
	bytes := []byte(str)

	for head, tail := 0, len(bytes)-1; head < tail; head, tail = head+1, tail-1 {
		bytes[head], bytes[tail] = bytes[tail], bytes[head]
	}

	return string(bytes)
}
