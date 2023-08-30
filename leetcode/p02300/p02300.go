package p02300

import "sort"

// Time Complexity: O(nlogn)
// Space Complexity: O(n)
func successfulPairs(spells []int, potions []int, success int64) []int {
	result := []int{}

	sort.Ints(potions)

	for _, n := range spells {
		left, right := 0, len(potions)-1

		for left <= right {
			mid := left + (right-left)/2
			if int64(potions[mid]*n) >= success {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}

		if left >= len(potions) {
			result = append(result, 0)
		} else {
			result = append(result, len(potions)-left)
		}
	}

	return result
}

/**
題目分析:
換言之就是用 binary search 找到第一個符合的

          i
          j
[ 1 2 3 4 5 ], *= 5, success = 25
當我找到3的時候我不能保證3以下還沒有沒有, 所以我要把 right 固定在 3 繼續找
但如果我發現 m 不符合的話就可以直接把他去掉 left = mid+1
*/
