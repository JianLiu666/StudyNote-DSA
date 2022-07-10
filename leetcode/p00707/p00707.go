package p00707

type Node struct {
	Val  int
	Next *Node
	Prev *Node
}

type MyLinkedList struct {
	head *Node
	tail *Node
	size int
}

func Constructor() MyLinkedList {
	return MyLinkedList{
		head: nil,
		tail: nil,
		size: 0,
	}
}

func (this *MyLinkedList) Get(index int) int {
	if this.size == 0 || index < 0 || index >= this.size {
		return -1
	}

	current := this.head
	for i := 0; i < index; i++ {
		current = current.Next
	}

	return current.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	node := &Node{Val: val, Next: nil, Prev: nil}
	if this.size == 0 {
		this.head = node
		this.tail = node
		this.size++
		return
	}

	node.Next = this.head
	this.head.Prev = node
	this.head = node
	this.size++
}

func (this *MyLinkedList) AddAtTail(val int) {
	node := &Node{Val: val, Next: nil, Prev: nil}
	if this.size == 0 {
		this.head = node
		this.tail = node
		this.size++
		return
	}

	node.Prev = this.tail
	this.tail.Next = node
	this.tail = node
	this.size++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.size {
		return
	}
	if index < 0 {
		index = 0
	}
	if index == 0 {
		this.AddAtHead(val)
		return
	}
	if index == this.size {
		this.AddAtTail(val)
		return
	}

	current := this.head
	for i := 0; i < index-1; i++ {
		current = current.Next
	}

	node := &Node{Val: val, Next: nil, Prev: nil}
	node.Next = current.Next
	node.Prev = current
	current.Next.Prev = node
	current.Next = node
	this.size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if this.size == 0 || index < 0 || index >= this.size {
		return
	}

	current := this.head
	for i := 0; i < index; i++ {
		current = current.Next
	}

	if current.Prev != nil {
		current.Prev.Next = current.Next
	} else {
		this.head = current.Next
	}

	if current.Next != nil {
		current.Next.Prev = current.Prev
	} else {
		this.tail = current.Prev
	}

	current.Next = nil
	current.Prev = nil
	this.size--
}
