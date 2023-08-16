package p00084

// Time Complexity: O(n)
// Space Complexity: O(n)
func largestRectangleArea_multipass(heights []int) int {
	result := 0

	// 初始化對每個 heights[i] 而言的左右邊界, 表示範圍內至少都是跟 heights[i] 一樣高或是比 heights[i] 更高的元素
	nextSmaller := make([]int, len(heights))
	prevSmaller := make([]int, len(heights))
	for i := 0; i < len(heights); i++ {
		nextSmaller[i] = len(heights)
		prevSmaller[i] = -1
	}

	stack := []int{}

	// 找出 heights[i] 對應的每個 nextSmaller[i]
	// 核心思路是用 monotonic stack 紀錄當前最高的位置, 只要遇到比 stack top 還低的元素, 就表示是 stack top 的 nextSmaller
	for i, h := range heights {
		for len(stack) > 0 && heights[stack[len(stack)-1]] > h {
			nextSmaller[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	stack = []int{}

	// 找出 heights[i] 對應的每個 prevSmaller[i]
	for i, h := range heights {
		for len(stack) > 0 && heights[stack[len(stack)-1]] > h {
			stack = stack[:len(stack)-1]
		}
		// 想法一樣, 只要比當前最高的 stack top 還低的話就退掉
		// 當退完之後 stack 裡面還有元素的話, 表示他一定是比現在的 heights[i] 還要小的元素, 就可以把他的位置記下來
		if len(stack) > 0 {
			prevSmaller[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	// 接著就只要在遍歷一次, 把每個 height[i] 的 prevSmaller 跟 nextSmaller 都抓出來計算一次範圍後找最大值即可
	for i, h := range heights {
		area := h * (nextSmaller[i] - prevSmaller[i] - 1)
		result = max(result, area)
	}

	return result
}
