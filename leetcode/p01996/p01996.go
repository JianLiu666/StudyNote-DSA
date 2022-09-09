package p01996

import "sort"

// Time Complexity: O(nlogn)
// Space Complexity: O(n)
func numberOfWeakCharacters(properties [][]int) int {
	// attack 由大到小排序
	// defense 由小到大排序
	sort.Slice(properties, func(i, j int) bool {
		if properties[i][0] > properties[j][0] {
			return true
		}
		if properties[i][0] == properties[j][0] {
			return properties[i][1] < properties[j][1]
		}
		return false
	})

	count := 0
	maxDef := properties[0][1]
	for i := 1; i < len(properties); i++ {
		if maxDef > properties[i][1] {
			count++
		} else {
			maxDef = properties[i][1]
		}
	}

	return count
}

func Max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
