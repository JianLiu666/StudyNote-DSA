package q00015

import "sort"

// Time Complexity: O(n^2)
// Space Complexity: O(1)
func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}

	sort.Ints(nums)

	result := make([][]int, 0)
	for idx, num := range nums {
		if num > 0 {
			break
		}

		if idx > 0 && nums[idx] == nums[idx-1] {
			continue
		}

		head := idx + 1
		tail := len(nums) - 1
		for head < tail {
			sum := nums[idx] + nums[head] + nums[tail]
			if sum < 0 {
				head++
			} else if sum > 0 {
				tail--
			} else {
				result = append(result, []int{nums[idx], nums[head], nums[tail]})
				head++

				for nums[head] == nums[head-1] && head < tail {
					head++
				}
			}
		}
	}

	return result
}
