package p00150

import "strconv"

// Time Complexity: O(n)
// Space Complexity: O(n)
func evalRPN(tokens []string) int {
	stack := CreateStack()
	for _, str := range tokens {
		if str == "+" || str == "-" || str == "*" || str == "/" {
			stack.Put(operate(str, stack.Pop(), stack.Pop()))
		} else {
			num, _ := strconv.Atoi(str)
			stack.Put(num)
		}
	}

	return stack.Pop()
}

func operate(symbol string, right, left int) int {
	if symbol == "+" {
		return left + right
	} else if symbol == "-" {
		return left - right
	} else if symbol == "*" {
		return left * right
	}

	return left / right
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

func (s *Stack) Put(str int) {
	s.arr = append(s.arr, str)
	s.size++
}

func (s *Stack) Pop() int {
	if s.IsEmpty() {
		return -1
	}

	num := s.arr[s.size-1]
	s.arr = s.arr[:s.size-1]
	s.size--

	return num
}
