package p00232

type MyQueue struct {
	input  Stack
	output Stack
}

func Constructor() MyQueue {
	return MyQueue{
		input:  CreateStack(),
		output: CreateStack(),
	}
}

func (this *MyQueue) Push(x int) {
	this.input.Push(x)
}

func (this *MyQueue) Pop() int {
	if this.Empty() {
		return -1
	}

	if this.output.IsEmpty() {
		for !this.input.IsEmpty() {
			this.output.Push(this.input.Pop())
		}
	}

	return this.output.Pop()
}

func (this *MyQueue) Peek() int {
	if this.Empty() {
		return -1
	}

	if this.output.IsEmpty() {
		for !this.input.IsEmpty() {
			this.output.Push(this.input.Pop())
		}
	}

	return this.output.Peek()
}

func (this *MyQueue) Empty() bool {
	return this.input.Size()+this.output.Size() == 0
}

type Stack struct {
	arr  []int
	size int
}

func CreateStack() Stack {
	return Stack{
		arr:  make([]int, 0),
		size: 0,
	}
}
func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack) Size() int {
	return s.size
}

func (s *Stack) Push(val int) {
	s.arr = append(s.arr, val)
	s.size++
}

func (s *Stack) Peek() int {
	if s.IsEmpty() {
		return -1
	}

	return s.arr[s.size-1]
}

func (s *Stack) Pop() int {
	if s.IsEmpty() {
		return -1
	}

	val := s.arr[s.size-1]
	s.arr = s.arr[:s.size-1]
	s.size--

	return val
}
