package q00001

// Time Complexity: O(n)
// Space Complexity: O(n)
func twoSum(nums []int, target int) []int {
	seen := make(map[int]int, len(nums))

	for i, val := range nums {
		if j, exists := seen[target-val]; exists {
			return []int{j, i}
		}

		seen[val] = i
	}

	return nil
}
