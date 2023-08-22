package p00295

import "container/heap"

// 為了不要重複寫太多相同的程式碼, 用一個 struct 維護 priority queue 的資料結構
// 在建構時傳入 compare function 就好
type CustomHeap struct {
	Elements []int
	LessFunc func(i, j int) bool
}

func (h *CustomHeap) Len() int           { return len(h.Elements) }
func (h *CustomHeap) Peek() int          { return h.Elements[0] }
func (h *CustomHeap) Swap(i, j int)      { h.Elements[i], h.Elements[j] = h.Elements[j], h.Elements[i] }
func (h *CustomHeap) Less(i, j int) bool { return h.LessFunc(h.Elements[i], h.Elements[j]) }
func (h *CustomHeap) Push(v any)         { h.Elements = append(h.Elements, (v.(int))) }
func (h *CustomHeap) Pop() (v any) {
	v, h.Elements = h.Elements[h.Len()-1], h.Elements[:h.Len()-1]
	return
}

type MedianFinder struct {
	smallHeap *CustomHeap // max-heap, 因為實際存放的資料是較小的那部分, 所以命名成 small heap
	largeHeap *CustomHeap // min-heap, 同上, 反之就是 large heap
}

func Constructor() MedianFinder {
	return MedianFinder{
		smallHeap: &CustomHeap{
			// max heap compare function
			LessFunc: func(i, j int) bool { return i > j },
		},
		largeHeap: &CustomHeap{
			// min heap compare function
			LessFunc: func(i, j int) bool { return i < j },
		},
	}
}

func (this *MedianFinder) AddNum(num int) {
	// 同時維護兩個 heaps 的原因
	// 為了能在 find median 時能馬上取出中間值, 所以我們需要平均擺放元素進兩個 heap
	// min-heap 需要維護較大的一半, 取出來的 minimum 才會是最靠近中間值的元素
	// 同上, max-heap 維護較小的一半, 這樣取出 maximum 時才會是最靠近中間的元素
	// 因此現在要討論的就只有如何將元素各自歸類在他們該去的那個 heaps
	//
	// 我優先將元素歸類在較大的那半(large(min-heap)), 加進去後要把目前最小的取出來放進較小的那半(small(max-heap))
	// 這樣就可以保證現在的 max-heap 裡面最大的元素是我在 min-heap 裡面最小的, 然後我現在的 min-heap 又有一個最小的
	// 這時候一但需要 find median 時就會是兩個最鄰近的元素在取中間值
	// 當陣列元素總數量在奇數時反之亦然
	if (this.smallHeap.Len()+this.largeHeap.Len())%2 == 0 {
		heap.Push(this.largeHeap, num)
		heap.Push(this.smallHeap, heap.Pop(this.largeHeap))
	} else {
		heap.Push(this.smallHeap, num)
		heap.Push(this.largeHeap, heap.Pop(this.smallHeap))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if (this.smallHeap.Len()+this.largeHeap.Len())%2 == 0 {
		return float64(this.smallHeap.Peek()+this.largeHeap.Peek()) / 2
	} else {
		return float64(this.smallHeap.Peek())
	}
}
