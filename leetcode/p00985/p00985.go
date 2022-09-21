package p00985

// Time Complexity: O(n)
// Space Complexity: O(1)
func sumEvenAfterQueries(nums []int, queries [][]int) []int {
	evenSum := 0
	for _, n := range nums {
		if n&1 == 0 {
			evenSum += n
		}
	}

	res := []int{}

	for _, query := range queries {
		if nums[query[1]]&1 == 0 {
			evenSum -= nums[query[1]]
		}

		nums[query[1]] += query[0]
		if nums[query[1]]&1 == 0 {
			evenSum += nums[query[1]]
		}

		res = append(res, evenSum)
	}

	return res
}
