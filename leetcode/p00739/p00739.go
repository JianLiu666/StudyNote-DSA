package p00739

// Time Complexity: O(n)
// Space Complexity: O(n)
func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))

	stack := CreateStack()
	for cur, temp := range temperatures {
		for !stack.IsEmpty() && temp > temperatures[stack.Peek()] {
			res[stack.Peek()] = cur - stack.Peek()
			stack.Pop()
		}
		stack.Put(cur)
	}
	return res
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

func (s *Stack) Put(val int) {
	s.arr = append(s.arr, val)
	s.size++
}

func (s *Stack) Pop() {
	if s.IsEmpty() {
		return
	}

	s.arr = s.arr[:s.size-1]
	s.size--
}

func (s *Stack) Peek() int {
	return s.arr[s.size-1]
}
