package p00334

func increasingTriplet_monotonic(nums []int) bool {
	q := []int{}

	for _, n := range nums {
		if len(q) == 0 || q[len(q)-1] < n {
			// 只要符合單調遞增, 就直接往後累加
			q = append(q, n)
		} else {
			// 這邊就是單調遞增在這題應用上最巧妙的地方, 只要不能往後塞, 那就把前面的一個數字替換掉
			// 這邊可能會有的疑慮是替換之後到底還符不符合題目說的 i<j<k
			// 直接看範例:
			// nums = [20, 100, 10, 12, 5, 13]
			// 在上述這個情境中 q 裡面的 element 是 [20, 100], 直到遇到 10 之後會替換成 [10, 100]
			// 雖然 queue 裡面的元素被替換了, 但不代表會影響到排列的順序, 如果接下來遇到的是 100 以上的元素, 不管是 [20, 100] 或 [10, 100] 都能符合
			// 繼續按照題目往下走遇到 12, 此時變成 [10, 12] 也就是說當兩次都替換掉之後就又回的 i < j 的順序了
			idx := 0
			for idx < len(q) && q[idx] < n {
				idx++
			}
			q[idx] = n
		}
		if len(q) >= 3 {
			return true
		}
	}

	return len(q) >= 3
}

// 20, 100, 10, 12, 5, 13

// small 20  10  5
// big   100 12
