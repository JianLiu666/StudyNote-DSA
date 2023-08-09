package p00525

// Time Complexity: O(n)
// Space Complexity: O(n)
func findMaxLength(nums []int) int {
	result := 0

	// 用來記錄 presum 的過程
	// key: 總和
	// value: 最後出現位置
	memo := map[int]int{
		0: -1, // init
	}

	presum := 0
	for i, n := range nums {
		// 把題目改變一下, 0,1 -> -1,1
		// 目的是更方便的計算總和, 就可以使用 prefix sum 的技巧
		if n == 1 {
			presum += 1
		} else {
			presum -= 1
		}

		// 當曾經出現過一樣的累加結果時就表示, i-j 位置中間的計數總和一定為0
		// e.g. prefix sum = 2
		// j                i
		// 11               00000011111111
		// 0111
		// 001111
		//
		// 不管怎麼相減都可以保證最後剩下的 0,1 個數相等, 符合題目要求
		if j, exists := memo[presum]; exists {
			result = max(result, i-j)
		} else {
			memo[presum] = i
		}
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
