package p00451

import "strings"

type Character struct {
	Val  rune
	Freq int
}

// Time Complexity: O(nlogn)
// Space Complexity: O(n)
func frequencySort(s string) string {
	// 紀錄每個 character 的出現頻率
	memo := map[rune]*Character{}
	for _, ch := range s {
		if _, exists := memo[ch]; exists {
			memo[ch].Freq++
		} else {
			memo[ch] = &Character{
				Val:  ch,
				Freq: 1,
			}
		}
	}

	maxHeap := make([]*Character, len(memo))
	cur, end := len(memo)-1, len(memo)-1

	for _, info := range memo {
		maxHeap[cur] = info
		shiftDown(maxHeap, cur, end)
		cur--
	}

	var str strings.Builder
	for i := 0; i < len(maxHeap); i++ {
		for j := 0; j < maxHeap[0].Freq; j++ {
			str.WriteRune(maxHeap[0].Val)
		}

		maxHeap[0], maxHeap[end] = maxHeap[end], maxHeap[0]
		end--
		shiftDown(maxHeap, 0, end)
	}

	return str.String()
}

func shiftDown(heap []*Character, start, end int) {
	parent := start
	for {
		child := parent*2 + 1
		if child > end {
			break
		}

		if child+1 <= end && heap[child+1].Freq > heap[child].Freq {
			child++
		}

		if heap[child].Freq > heap[parent].Freq {
			heap[child], heap[parent] = heap[parent], heap[child]
			parent = child
		} else {
			break
		}
	}
}
