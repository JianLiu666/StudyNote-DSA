package p00621

import "container/heap"

// Time Complexity: O(nlogn)
// Space Complexity: O(n)
func leastInterval(tasks []byte, n int) int {
	// 用來記錄冷卻時間
	memo := map[byte]int{}
	pq := &PriorityQueue{}
	heap.Init(pq)

	// 按照間隔順序加入 tasks
	for _, task := range tasks {
		heap.Push(pq, &Task{
			name: task,
			time: memo[task],
		})
		memo[task] += n + 1
	}

	time := 0
	for pq.Len() > 0 {
		if (*pq)[0].time <= time {
			heap.Pop(pq)
		}
		time++
	}

	return time
}

type Task struct {
	name byte
	time int
}

type PriorityQueue []*Task

func (q PriorityQueue) Len() int           { return len(q) }
func (q PriorityQueue) Swap(i, j int)      { q[i], q[j] = q[j], q[i] }
func (q PriorityQueue) Less(i, j int) bool { return q[i].time < q[j].time }
func (q *PriorityQueue) Push(v any)        { *q = append(*q, v.(*Task)) }
func (q *PriorityQueue) Pop() any {
	result := (*q)[len(*q)-1]
	*q = (*q)[:len(*q)-1]
	return result
}
