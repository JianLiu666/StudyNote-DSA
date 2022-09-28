package p00055

// Time Complexity: O(n)
// Space Complexity: O(n)
func canJump_bottomUp(nums []int) bool {
	arrived := make([]bool, len(nums))
	arrived[0] = true

	for i, n := range nums {
		// 一旦發現有抵達不了的位置, 那就不用在繼續了
		if !arrived[i] {
			return false
		}
		if i+n >= len(nums)-1 {
			return true
		}

		// 把這次能走到的最遠距離由後往前逐一標記
		for tail := i + n; tail > i; tail-- {
			// 只要連最後一格都標記過, 就不用再重複運算
			if arrived[tail] {
				break
			}
			arrived[tail] = true
		}
	}

	return false
}

// Time Complexity: O(n)
// Space Complexity: O(1)
func canJump_bottomUp_optimize(nums []int) bool {
	arrived := 0
	for i, n := range nums {
		// 一旦發現有抵達不了的位置, 那就不用在繼續了
		if arrived < i {
			return false
		}
		// 只要能夠抵達終點, 就提早結束
		if arrived >= len(nums)-1 {
			return true
		}
		arrived = max(arrived, i+n)
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
