package p00239

func maxSlidingWindow(nums []int, k int) []int {
	result := []int{}

	// 用 slice 模擬 monotonic queue
	queue := []int{}

	for i := 0; i < len(nums); i++ {
		// 確保 queue 裡面的 elemnts 以 desc 排序
		for len(queue) > 0 && nums[queue[len(queue)-1]] < nums[i] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, i)

		// 確保 queue 的 element indexes 不會超出 sliding window 的範圍
		for len(queue) > 0 && queue[0] <= i-k {
			queue = queue[1:]
		}

		// 每次滑動 sliding window 時都從 queue 裡面取出最大值
		if i >= k-1 {
			result = append(result, nums[queue[0]])
		}

	}

	return result
}
