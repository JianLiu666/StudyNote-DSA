package p00134

// Time Complexity: O(n)
// Space Complexity: O(1)
func canCompleteCircuit(gas []int, cost []int) int {
	totalSurplus := 0 // 記錄整趟旅程中的汽油總盈餘
	surplus := 0      // 紀錄 "目前為止" 的汽油盈餘
	start := 0        // 紀錄建議起始點

	for i := 0; i < len(gas); i++ {
		totalSurplus += gas[i] - cost[i]
		surplus += gas[i] - cost[i]
		if surplus < 0 {
			// 當 surplus 出現負數時就代表從原本的 start 無法走到現在這個 station
			// 因此是一個不好的起點且中間也不會是好的起點, 因此我們將起點放在下一個 station 開始
			surplus = 0
			start = i + 1
		}
	}

	// 如果單純計算盈餘都無法為正時, 代表任何一個起始點都不可能成功
	if totalSurplus < 0 {
		return -1
	}
	return start
}
