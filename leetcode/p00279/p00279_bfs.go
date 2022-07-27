package p00279

import (
	"math"
)

// Time Complexity: O(n)
// Space Complexity: O(n)
func numSquares(n int) int {
	if n <= 3 {
		return n
	}

	q := CreateUnlimitQueue()
	q.Enqueue(n, 1)

	step := 0
	for !q.IsEmpty() {
		size := q.Size()
		step++

		for i := 0; i < size; i++ {
			data := q.Dequeue()

			divisor := int(math.Sqrt(float64(data.Target)))
			square := divisor * divisor
			for divisor >= data.MinNum {
				if data.Target == square {
					return step
				}

				q.Enqueue(data.Target-square, divisor)
				divisor--
				square = divisor * divisor
			}

			data = nil
		}
	}

	return -1
}

type Node struct {
	Target int
	MinNum int
	Next   *Node
	Prev   *Node
}

func CreateNode(target, minNum int) *Node {
	return &Node{
		Target: target,
		MinNum: minNum,
		Next:   nil,
		Prev:   nil,
	}
}

type UnlimitQueue struct {
	head *Node
	tail *Node
	size int
}

func CreateUnlimitQueue() UnlimitQueue {
	return UnlimitQueue{
		head: nil,
		tail: nil,
		size: 0,
	}
}

func (this *UnlimitQueue) IsEmpty() bool {
	return this.size == 0
}

func (this *UnlimitQueue) Size() int {
	return this.size
}

func (this *UnlimitQueue) Enqueue(target, minNum int) {
	node := CreateNode(target, minNum)

	if this.head == nil {
		this.head = node
		this.tail = node
	} else {
		this.tail.Next = node
		node.Prev = this.tail
		this.tail = node
	}

	this.size++
}

func (this *UnlimitQueue) Dequeue() *Node {
	if this.IsEmpty() {
		return nil
	}

	node := this.head
	if this.size == 1 {
		this.head = nil
		this.tail = nil
	} else {
		this.head = this.head.Next
		this.head.Prev = nil
	}
	node.Next = nil
	node.Prev = nil

	this.size--
	return node
}
