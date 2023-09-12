package p00918

import "math"

// Time Complexity: O(n+k), where k is the length of words2
// Space Complexity: O(1)
func maxSubarraySumCircular(nums []int) int {
	curmax, maximum, curmin, minimum, total := 0, math.MinInt, 0, math.MaxInt, 0

	for _, n := range nums {
		curmax = max(curmax+n, n)
		maximum = max(maximum, curmax)
		curmin = min(curmin+n, n)
		minimum = min(minimum, curmin)
		total += n
	}

	if total == curmin {
		return maximum
	}
	return max(maximum, total-minimum)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/**
分析問題:
因為要考慮頭尾循環的關係, 所以可能的答案有兩種情況:

case1. [x x x o o o o o o o x x x x] 或
case2. [o o o x x x x x x x o o o o] o = 我們想要的 subarray

case1 可以直接用 dp 求出
case2 可以以換個方式思考, 如果我們想要求出外側的最大值, 我們可以轉向找最小的 minarray, 用 sum(arr) - minsubarray 就是外側的總和

最後我們只要比較 maximum subarray 與 sum(arr) - minimum subarray 誰比較大就是答案了

必須注意一件事情是, 如果當 array 裡面的元素都是負值時, sum(arr) 與 minsubarray 找到的結果都會是一樣的, 此時相減後變成是0
因此我們必須直接返回 maximum subarray 的結果
*/
