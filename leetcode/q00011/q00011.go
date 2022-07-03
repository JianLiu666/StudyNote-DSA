package q00011

// Time Complexity: O(n)
// Space Complexity: O(1)
func maxArea(height []int) int {
	head := 0
	tail := len(height) - 1

	max_area := 0
	for head != tail {
		w := tail - head
		h := height[head]
		if h > height[tail] {
			h = height[tail]
		}

		area := h * w
		if area > max_area {
			max_area = area
		}

		if height[head] < height[tail] {
			head++
		} else {
			tail--
		}
	}

	return max_area
}
