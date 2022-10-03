package p01578

// Time Complexity: O(n)
// Space Complexity: O(1)
func minCost(colors string, neededTime []int) int {
	color := colors[0]
	cost := neededTime[0]

	res := 0

	for i := 1; i < len(colors); i++ {
		if colors[i] == color {
			// 遇到相同顏色就比較處理時間, 選擇較小的成本, 將大的記錄下來準備下次比較
			if cost < neededTime[i] {
				res += cost
				cost = neededTime[i]
			} else {
				res += neededTime[i]
			}
		} else {
			// 紀錄當下的氣球顏色與處理時間
			color = colors[i]
			cost = neededTime[i]
		}
	}

	return res
}
