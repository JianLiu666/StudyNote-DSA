package p00164

import "math"

// Time Complexity: O(n), where n is the number of nums
// Space Complexity: O(n)
func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}

	// 找出最大數字, 用來決定後續排序所需比較的位數
	maximum := math.MinInt
	for _, n := range nums {
		if n > maximum {
			maximum = n
		}
	}

	// 從小到大對每個位數排序一遍 (least significant digit)
	exp := 1
	for maximum/exp > 0 {
		buckets := [10]int{}
		tmp := make([]int, len(nums))
		// radixing
		for _, n := range nums {
			buckets[(n/exp)%10]++
		}
		// indexing
		for i := 1; i < 10; i++ {
			buckets[i] += buckets[i-1]
		}
		// sorting
		for i := len(nums) - 1; i >= 0; i-- {
			idx := (nums[i] / exp) % 10
			buckets[idx]--
			tmp[buckets[idx]] = nums[i]
		}
		nums = tmp
		exp *= 10
	}

	// 找出相鄰兩數的最大差異
	ans := math.MinInt
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] > ans {
			ans = nums[i] - nums[i-1]
		}
	}
	return ans
}
