package p00300

// Time Complexity: O(nlogn)
// Space Complexity: O(n)
func lengthOfLIS_binarysearch(nums []int) int {
	seq := []int{nums[0]}

	tailIdx := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > seq[tailIdx] {
			seq = append(seq, nums[i])
			tailIdx++
		} else if nums[i] < seq[tailIdx] {
			replaceElementInArray(seq, nums[i])
		}
	}

	return len(seq)
}

func replaceElementInArray(seq []int, target int) {
	head := 0
	tail := len(seq) - 1
	for head < tail {
		mid := head + (tail-head)/2
		if seq[mid] < target {
			head = mid + 1
		} else {
			tail = mid
		}
	}
	seq[head] = target
}

// Time Complexity: O(n^2)
// Space Complexity: O(n)
func lengthOfLIS_dp(nums []int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}

	max := 1
	for current := 0; current < len(dp); current++ {
		for compare := 0; compare < current; compare++ {
			if nums[current] > nums[compare] && dp[current] < dp[compare]+1 {
				dp[current] = dp[compare] + 1
				if max < dp[current] {
					max = dp[current]
				}
			}
		}
	}

	return max
}
