package p00042

func trap(height []int) int {
	result := 0

	// monotonic stack (desc)
	stack := []int{}

	for i := 0; i < len(height); i++ {
		// 當遇到右邊比 stack top 還高的時候, 代表這個區間可以蓄水
		for len(stack) > 0 && height[i] >= height[stack[len(stack)-1]] {
			base := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// 找出地基高度後沒有左邊的牆就不可能蓄水了
			if len(stack) == 0 {
				break
			}
			// 取出兩邊最低的高度後乘上寬度
			h := min(height[stack[len(stack)-1]], height[i])
			result += (h - height[base]) * (i - stack[len(stack)-1] - 1)
		}
		stack = append(stack, i)
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
