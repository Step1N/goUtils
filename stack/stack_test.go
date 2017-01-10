package stack

import (
	"testing"
)

func TestPushElementInStack(t *testing.T) {
	s := NewStack()
	s.Push("a")
	s.Push("b")
	s.Push("d")
	s.Push("e")
}

func TestPopElementInStack(t *testing.T) {
	s := NewStack()
	s.Push("a")
	s.Push("b")
	s.Push("d")
	s.Push("e")
	s.Push("f")
	s.Pop()
	s.Pop()
	s.Pop()
}

func TestPeekElementInStack(t *testing.T) {
	s := NewStack()
	s.Push("a")
	s.Push("b")
	s.Push("d")
	s.Push("e")
	s.Push("f")
}
