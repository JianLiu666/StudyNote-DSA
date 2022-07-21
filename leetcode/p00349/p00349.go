package p00349

// Time Complexity: O(n)
// Space Complexity: O(n)
func intersection(nums1 []int, nums2 []int) []int {
	nums := nums1
	stream := nums2
	if len(nums2) < len(nums1) {
		nums = nums2
		stream = nums1
	}

	table := map[int]bool{}
	for _, n := range nums {
		table[n] = true
	}

	res := []int{}
	for i := 0; i < len(stream); i++ {
		if unvisited, exists := table[stream[i]]; exists && unvisited {
			res = append(res, stream[i])
			table[stream[i]] = false
		}
	}

	return res
}

// Time Complexity: O(n)
// Space Complexity: O(n)
func intersection_sol2(nums1 []int, nums2 []int) []int {
	table1 := map[int]int{}
	table2 := map[int]int{}

	for _, n := range nums1 {
		table1[n]++
	}
	for _, n := range nums2 {
		table2[n]++
	}

	res := []int{}
	for k := range table1 {
		if _, exists := table2[k]; exists {
			res = append(res, k)
		}
	}

	return res
}
