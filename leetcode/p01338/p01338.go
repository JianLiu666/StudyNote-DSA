package p01338

import "sort"

// Time Complexity: O(n)
// Space Complexity: O(n)
func minSetSize_hash(arr []int) int {
	occurs := map[int]int{}
	for idx, n := range arr {
		occurs[n]++
		arr[idx] = 0
	}

	for _, val := range occurs {
		arr[val-1]++
	}

	goal := len(arr) / 2
	cnt := 0
	for cur := len(arr) - 1; cur >= 0 && goal > 0; cur-- {
		if arr[cur] > 0 {
			for i := 0; i < arr[cur] && goal > 0; i++ {
				goal -= cur + 1
				cnt++
			}
		}
	}

	return cnt
}

// Time Complexity: O(nlogn)
// Space Complexity: O(n)
func minSetSize_sorting(arr []int) int {
	occurs := map[int]int{}
	cursor := 0
	for idx, n := range arr {
		occurs[n]++
		if occurs[n] == 1 {
			arr[cursor], arr[idx] = arr[idx], arr[cursor]
			cursor++
		}
	}

	sort.Slice(arr[:cursor], func(i, j int) bool {
		return occurs[arr[i]] < occurs[arr[j]]
	})

	goal := len(arr) / 2
	cnt := 0
	for goal > 0 {
		goal -= occurs[arr[cursor-1]]
		cursor--
		cnt++
	}

	return cnt
}
