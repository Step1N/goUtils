package stack


type Stack struct {
	top  *Frame
	size int
}
type Frame struct {
	value interface{}
	next  *Frame
}

func NewStack() *Stack {
	s := new(Stack)
	return s
}

func (s *Stack) Length() int {
	return s.size
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack) Push(val interface{}) {
	s.top = &Frame{val, s.top}
	s.size++
}

func (s *Stack) Peek() interface{} {
	return s.top.value
}

func (s *Stack) Pop() (val interface{}) {
	if s.size > 0 {
		val, s.top = s.top.value, s.top.next
		s.size--
		return
	}

	return -1
}
