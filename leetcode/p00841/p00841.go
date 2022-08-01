package p00841

// Time Complexity: O(n)
// Space Complexity: O(n)
func canVisitAllRooms(rooms [][]int) bool {
	visited := make([]bool, len(rooms))
	visited[0] = true
	visitedCount := 1

	q := CreateQueue()
	visitedCount += visitRoom(&q, visited, rooms, 0)

	for !q.Empty() {
		roomId := q.Pop()
		visitedCount += visitRoom(&q, visited, rooms, roomId)
	}

	return visitedCount == len(rooms)
}

func visitRoom(q *queue, visited []bool, rooms [][]int, roomId int) int {
	visitedCount := 0
	for _, key := range rooms[roomId] {
		if !visited[key] {
			visited[key] = true
			visitedCount++
			q.Put(key)
		}
	}

	return visitedCount
}

type queuekNode struct {
	Val  int
	Next *queuekNode
	Prev *queuekNode
}

func CreateQueueNode(val int) *queuekNode {
	return &queuekNode{
		Val: val,
	}
}

type queue struct {
	head *queuekNode
	tail *queuekNode
	size int
}

func CreateQueue() queue {
	return queue{
		head: nil,
		tail: nil,
		size: 0,
	}
}

func (q *queue) Empty() bool {
	return q.size == 0
}

func (q *queue) Size() int {
	return q.size
}

func (q *queue) Put(val int) {
	node := CreateQueueNode(val)

	if q.size == 0 {
		q.head = node
		q.tail = node
	} else {
		node.Next = q.head
		q.head.Prev = node
		q.head = node
	}
	q.size++
}

func (q *queue) Pop() int {
	if q.size == 0 {
		return -1
	}

	node := q.tail
	if q.size == 1 {
		q.head = nil
		q.tail = nil
	} else {
		q.tail = q.tail.Prev
		q.tail.Next = nil
	}
	node.Next = nil
	node.Prev = nil
	q.size--

	return node.Val
}
