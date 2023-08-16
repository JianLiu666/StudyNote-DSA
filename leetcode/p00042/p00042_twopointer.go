package p00042

func trap_twopointer(height []int) int {
	result := 0

	mostLeft := make([]int, len(height))
	mostRight := make([]int, len(height))

	for i := 1; i < len(height); i++ {
		mostLeft[i] = max(mostLeft[i-1], height[i-1])
	}
	for i := len(height) - 2; i >= 0; i-- {
		mostRight[i] = max(mostRight[i+1], height[i+1])
	}

	for i := 0; i < len(height); i++ {
		h := min(mostLeft[i], mostRight[i]) - height[i]
		result += max(0, h*1)
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
