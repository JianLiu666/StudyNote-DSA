package p00128

// Time Complexity: O()
// Space Complexity: O()
func longestConsecutive(nums []int) int {
	memo := map[int]bool{}
	for _, n := range nums {
		memo[n] = true
	}

	maximum := 0
	for n := range memo {
		// 此時表示 n 為起點
		if !memo[n-1] {
			incr := n
			count := 0
			for memo[incr] {
				incr++
				count++
			}

			if count > maximum {
				maximum = count
			}
		}
	}

	return maximum
}
