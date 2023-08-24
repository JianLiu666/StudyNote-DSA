package p00031

import "sort"

func nextPermutation(nums []int) {
	// 從後往前找到第一個非遞增的位置
	i := len(nums) - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	// 如果已經是最大數字的話, 直接做一次 ASC 排序後結束
	if i == -1 {
		sort.Ints(nums)
		return
	}

	// 找到第一個比定位點數字還要大的數字後交換, 再把定位點後一個以後的數字重新以 ASC 排序過一遍
	j := len(nums) - 1
	for j > i && nums[i] >= nums[j] {
		j--
	}
	nums[i], nums[j] = nums[j], nums[i]
	sort.Ints(nums[i+1:])
}
