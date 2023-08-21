package p00787

import "container/heap"

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	// 轉換成 graph
	cities := make([][][2]int, n)
	for _, tuple := range flights {
		from, to, price := tuple[0], tuple[1], tuple[2]
		cities[from] = append(cities[from], [2]int{to, price})
	}

	// 用 builtin heap 實作 priority queue (a.k.a min-heap)
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Travel{to: src, cost: 0, times: 0})

	// 用來記錄走過的路線, 節省時間
	// [CITY_INDEX][FLIGHT_INDEX]bool
	// 題目給定的 k 代表的是 "轉機次數", 所以要把第一次直飛也算成一次, 但因為是從 1 開始數的所以空間必須要再多開 1 個
	visited := make([][]bool, n)
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, k+2)
	}

	for pq.Len() > 0 {
		travel := heap.Pop(pq).(*Travel)

		if travel.to == dst {
			return travel.cost
		}
		if travel.times == k+1 {
			continue
		}
		if visited[travel.to][travel.times] {
			continue
		}
		visited[travel.to][travel.times] = true

		for _, ticket := range cities[travel.to] {
			if visited[ticket[0]][travel.times+1] {
				continue
			}
			heap.Push(pq, &Travel{
				to:    ticket[0],
				cost:  travel.cost + ticket[1],
				times: travel.times + 1,
			})
		}
	}

	return -1
}

type Travel struct {
	to    int // 抵達城市
	cost  int // 目前總花費
	times int // 目前總飛行次數
}

type PriorityQueue []*Travel

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].cost < pq[j].cost }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(v any)        { *pq = append(*pq, v.(*Travel)) }
func (pq *PriorityQueue) Pop() any {
	result := (*pq)[len(*pq)-1]
	*pq = (*pq)[:len(*pq)-1]
	return result
}
