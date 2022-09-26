package p02343

// Time Complexity: O(mn), where m is the number of nums, n is the length of each num
// Space Complexity: O(mn)
func smallestTrimmedNumbers(nums []string, queries [][]int) []int {
	type Node struct {
		num string
		idx int
	}

	length := len(nums[0])

	// 用來記錄 radix sort 的每一個過程
	memo := map[int][]Node{}
	radix := []Node{}

	for i, num := range nums {
		radix = append(radix, Node{num, i})
	}

	// LSD Radix Sort
	for i := 0; i < length; i++ {
		// counting
		buckets := [10][]Node{}
		idx := length - 1 - i
		for _, node := range radix {
			buckets[node.num[idx]-'0'] = append(buckets[node.num[idx]-'0'], node)
		}
		// sort
		radix = []Node{}
		for _, bucket := range buckets {
			for _, node := range bucket {
				radix = append(radix, node)
			}
		}
		memo[i+1] = radix
	}

	ans := []int{}
	for _, query := range queries {
		ans = append(ans, memo[query[1]][query[0]-1].idx)
	}
	return ans
}
