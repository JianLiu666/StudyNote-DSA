package q00136

// Time Complexity: O(n)
// Space Complexity: O(1)
func singleNumber(nums []int) int {
	answer := 0
	for _, num := range nums {
		answer ^= num
	}

	return answer
}

// Time Complexity: O(n)
// Space Complexity: O(n)
func singleNumber_dsa(nums []int) int {
	seen := make(map[int]int, (len(nums)/2)+1)

	for _, num := range nums {
		seen[num] += 1
	}

	for num, count := range seen {
		if count == 1 {
			return num
		}
	}

	return -1
}
