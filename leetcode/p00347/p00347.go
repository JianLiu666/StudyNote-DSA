package p00347

// Time Complexity: O(n)
// Space Complexity: O(n)
func topKFrequent(nums []int, k int) []int {
	buckets := make([][]int, len(nums)+1)
	counter := map[int]int{}

	for _, num := range nums {
		counter[num]++
	}

	for num, count := range counter {
		if buckets[count] == nil {
			buckets[count] = []int{}
		}
		buckets[count] = append(buckets[count], num)
	}

	res := []int{}
	cursor := 0
	for i := len(buckets) - 1; i > -1 && cursor < k; i-- {
		if len(buckets[i]) > 0 {
			res = append(res, buckets[i]...)
			cursor += len(buckets[i])
		}
	}

	return res
}
