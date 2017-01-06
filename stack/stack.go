package stack

import "fmt"

type Stack struct {
	input []interface{}
	count int
}

func NewStack() *Stack {
	s := newStack()
	return s
}

func newStack() *Stack {
	in := make([]interface{}, 0)
	return &Stack{input: in, count: -1}
}

func (s *Stack) Push(element interface{}) {
	if s.count == -1 {
		ns := newStack()
		ns.input = append(ns.input, element)
		s.input = ns.input
		s.count = s.count + 1
		return
	}

	s.input = append(s.input, element)
}

func (s *Stack) Pop() interface{} {
	if len(s.input) == 0 {
		s.count = -1
		return -1
	}

	el := s.input[s.count]
	s.input = s.input[s.count+1:]
	return el
}

func (s *Stack) Peek() interface{} {
	if len(s.input) == 0 {
		return -1
	}

	return s.input[s.count]
}

func (s *Stack) Size() int{

	return len(s.input)
}

func (s *Stack) Print() {
	if s.count == -1 {
		return
	}
	for i := 0; i < len(s.input); i++ {
		fmt.Printf("%v \t", s.input[i])
	}
}
