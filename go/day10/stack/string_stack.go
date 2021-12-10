package stack

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

func (s *StringStack) Peek() (string, bool) {
	if len(s.stack) == 0 {
		return "", false
	}
	return s.stack[len(s.stack)-1], true
}

func (s *StringStack) String() string {
	str := "BOTTOM |  "
	for _, r := range(s.stack) {
		sr := string(r)
		str += sr + " "
	}
	str += " | TOP"
	return str
}
