package p01143

// Time Complexity: O(mn), where m is the length of text1, n is the length of text2
// Space Complexity: O(mn)
func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([][]int, len(text1)+1)
	for i := range dp {
		dp[i] = make([]int, len(text2)+1)
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[len(text1)][len(text2)]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/**
分析問題:

      a b c d e
    0 1 2 3 4 5 6 7 8 9 0 text1
  0 0 0 0 0 0 0
a 1 0 1 1 1 1 1
c 2 0 1 1 2 2 2
e 3 0 1 1 2 2 3
  4
  5
  6
  7
  8
  9
  0
  text2

*/
