package stack

import "fmt"

type StringStack struct {
	stack []string
}

func NewStringStack() *StringStack {
	return &StringStack{
		stack: []string{},
	}
}

func (s *StringStack) Push(item string)  {
	s.stack = append(s.stack, item)
}

func (s *StringStack) Pop() (string, bool) {
	l := len(s.stack)
	if l == 0 {
		return "", false
	}
	popped := s.stack[l-1]
	s.stack = s.stack[:l-1]
	return popped, true
}

func (s *StringStack) Peek() string {
	return s.stack[len(s.stack)-1]
}

func (s *StringStack) String() string {
	return fmt.Sprintf("BOTTOM |  %v  | TOP", s.stack)
}
