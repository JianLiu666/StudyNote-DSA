package p00452

import "sort"

// Time Complexity: O(nlogn)
// Space Complexity: O(1)
func findMinArrowShots(points [][]int) int {
	result := 1

	// tc: O(nlogn)
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	// 拿第一個線段的 end 作為初始值
	end := points[0][1]
	for i := 1; i < len(points); i++ {
		if points[i][0] <= end {
			// 每次都先確定現在這個線段的 start 還在上一個 end 的範圍內
			// 然後再檢查 end 目前最小到哪裡
			end = min(end, points[i][1])
		} else {
			// 如果 start 已經超過 end 了, 代表這屬於第二個群集 (i.e., 需要再多一標)
			result++
			end = points[i][1]
		}
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/**
題目分析：
找出最大共同重疊區間
start 排序後, 一直合併 start[i] < end[i-1] 的

0    5    10   15   20        0    5    10   15   20
+----+----+----+----+         +----+----+----+----+
          -------              ------
   ------                ->     ------- i=1
 ------                              ------
        -----                           -------

過程中要一直找出最小的 end 位置, 一旦下一個 start 超過 end 之後就要重新來過了

*/
