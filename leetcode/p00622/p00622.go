package p00622

type MyCircularQueue struct {
	queue    []int
	head     int
	tail     int
	size     int
	capacity int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		queue:    make([]int, k),
		head:     0,
		tail:     k - 1,
		size:     0,
		capacity: k,
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}

	this.tail++
	if this.tail == this.capacity {
		this.tail = 0
	}

	this.queue[this.tail] = value
	this.size++

	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}

	this.head++
	if this.head == this.capacity {
		this.head = 0
	}

	this.size--

	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}

	return this.queue[this.head]
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}

	return this.queue[this.tail]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.size == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return this.size == this.capacity
}
