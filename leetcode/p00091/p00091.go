package p00091

import "strconv"

func numDecodings(s string) int {
	if len(s) == 0 {
		return 0
	}

	mapping := map[string]bool{}
	genMapping(mapping)

	// 給定初始值
	s = "0" + s
	dp := make([]int, len(s))
	dp[0] = 1
	if _, exists := mapping[s[1:2]]; exists {
		dp[1] = 1
	}

	for i := 2; i < len(s); i++ {
		if _, exists := mapping[s[i:i+1]]; exists {
			dp[i] += dp[i-1]
		}
		if _, exists := mapping[s[i-1:i+1]]; exists {
			dp[i] += dp[i-2]
		}
	}

	return dp[len(dp)-1]
}

func genMapping(m map[string]bool) {
	for i := 1; i <= 26; i++ {
		m[strconv.Itoa(i)] = true
	}
}
