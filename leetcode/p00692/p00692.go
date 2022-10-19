package p00692

import (
	"container/heap"
)

// Time Complexity: O(n+klogk), where n is the number of words, k is the number of unique words
// Space Complexity: O(n)
func topKFrequent(words []string, k int) []string {
	// tc: O(n)
	memo := map[string]int{}
	for _, word := range words {
		memo[word]++
	}

	// tc: O(n)
	h := &maxHeap{}
	for word, freq := range memo {
		h.Push(&info{
			word:      word,
			frequency: freq,
		})
	}

	// tc: O(klogk)
	heap.Init(h)
	res := []string{}
	for i := 0; i < k; i++ {
		res = append(res, heap.Pop(h).(*info).word)
	}

	return res
}

type info struct {
	word      string
	frequency int
}

type maxHeap []*info

func (h maxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h maxHeap) Len() int {
	return len(h)
}

func (h maxHeap) Less(i, j int) bool {
	if h[i].frequency > h[j].frequency {
		return true
	}
	if h[i].frequency == h[j].frequency {
		return h[i].word < h[j].word
	}
	return false
}

func (h *maxHeap) Push(v interface{}) {
	*h = append(*h, v.(*info))
}

func (h *maxHeap) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}
