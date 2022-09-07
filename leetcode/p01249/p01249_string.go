package p01249

import "strings"

// Time Complexity: O(n), where n is the length of s
// Space Complexity: O(n)
func minRemoveToMakeValid_string(s string) string {
	stack := CreateStringStack()

	var res strings.Builder
	for _, ch := range s {
		if ch != '(' && ch != ')' {
			if stack.Empty() {
				res.WriteRune(ch)
			} else {
				stack.Peek().WriteRune(ch)
			}

		} else if ch == '(' {
			// 遇到新的 '(' 就加入 stack
			var tmp strings.Builder
			tmp.WriteRune(ch)
			stack.Put(&tmp)

		} else if ch == ')' && !stack.Empty() {
			// 確定之前有遇過 '(' 才處理
			stack.Peek().WriteRune(ch)
			tmp := stack.Pop()
			if stack.Empty() {
				res.WriteString(tmp.String())
			} else {
				stack.Peek().WriteString(tmp.String())
			}
		}
	}

	// 把剩下只有 '(' 的字串處理完畢, 去掉 '(' 只保留字母
	for !stack.Empty() {
		tmp := stack.Pop().String()
		if stack.Empty() {
			res.WriteString(tmp[1:])
		} else {
			stack.Peek().WriteString(tmp[1:])
		}
	}

	return res.String()
}

type StringStack struct {
	arr  []*strings.Builder
	size int
}

func CreateStringStack() StringStack {
	return StringStack{
		arr:  []*strings.Builder{},
		size: 0,
	}
}

func (s *StringStack) Empty() bool {
	return s.size == 0
}

func (s *StringStack) Put(val *strings.Builder) {
	s.arr = append(s.arr, val)
	s.size++
}

func (s *StringStack) Peek() *strings.Builder {
	return s.arr[s.size-1]
}

func (s *StringStack) Pop() *strings.Builder {
	val := s.arr[s.size-1]
	s.size--
	s.arr = s.arr[:s.size]
	return val
}
