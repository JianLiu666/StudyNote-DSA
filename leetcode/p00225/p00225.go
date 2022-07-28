package p00225

type MyStack struct {
	queue Queue
}

func Constructor() MyStack {
	return MyStack{
		queue: CreateQueue(),
	}
}

func (this *MyStack) Push(x int) {
	size := this.queue.Size()
	this.queue.Push(x)
	for i := 0; i < size; i++ {
		this.queue.Push(this.queue.Pop())
	}
}

func (this *MyStack) Pop() int {
	return this.queue.Pop()
}

func (this *MyStack) Top() int {
	return this.queue.Peek()
}

func (this *MyStack) Empty() bool {
	return this.queue.Empty()
}

type Node struct {
	Val  int
	Next *Node
	Prev *Node
}

func CreateNode(val int) *Node {
	return &Node{
		Val:  val,
		Next: nil,
		Prev: nil,
	}
}

type Queue struct {
	head *Node
	tail *Node
	size int
}

func CreateQueue() Queue {
	return Queue{
		head: nil,
		tail: nil,
		size: 0,
	}
}

func (q *Queue) Size() int {
	return q.size
}

func (q *Queue) Empty() bool {
	return q.size == 0
}

func (q *Queue) Push(val int) {
	node := CreateNode(val)
	if q.head == nil {
		q.head = node
		q.tail = node
	} else {
		node.Next = q.head
		q.head.Prev = node
		q.head = node
	}

	q.size++
}

func (q *Queue) Pop() int {
	if q.Empty() {
		return -1
	}

	node := q.tail
	if q.size == 1 {
		q.head = nil
		q.tail = nil
	} else {
		q.tail = q.tail.Prev
		q.tail.Next = nil
		node.Prev = nil
	}

	q.size--

	return node.Val
}

func (q *Queue) Peek() int {
	if q.Empty() {
		return -1
	}

	return q.tail.Val
}
