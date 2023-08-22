package p00815

func numBusesToDestination(routes [][]int, source int, target int) int {
	result := 0

	// 紀錄每個站牌會經過的公車有哪些
	stop2bus := map[int][]int{}
	for busIdx, stopList := range routes {
		for _, stopIdx := range stopList {
			if _, exists := stop2bus[stopIdx]; !exists {
				stop2bus[stopIdx] = []int{}
			}
			stop2bus[stopIdx] = append(stop2bus[stopIdx], busIdx)
		}
	}

	// early pruning 條件
	busVisited := map[int]bool{}
	stopVisited := map[int]bool{}

	q := []int{source}
	stopVisited[source] = true

	for len(q) > 0 {
		size := len(q)
		for ; size > 0; size-- {
			stopIdx := q[0]
			q = q[1:]

			if stopIdx == target {
				return result
			}

			for _, busIdx := range stop2bus[stopIdx] {
				if busVisited[busIdx] {
					continue
				}
				busVisited[busIdx] = true
				for _, nextStopIdx := range routes[busIdx] {
					if stopVisited[nextStopIdx] {
						continue
					}
					stopVisited[nextStopIdx] = true
					q = append(q, nextStopIdx)
				}
			}
		}
		result++
	}

	return -1
}
