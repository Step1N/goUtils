package stack

import (
	"fmt"
	"testing"
)

func TestPushElementInStack(t *testing.T) {
	s := newStack()
	s.Push("a")
	s.Push("b")
	s.Push("d")
	s.Push("e")
	fmt.Println("Stack elements")
	s.Print()
}

func TestPopElementInStack(t *testing.T) {
	s := newStack()
	s.Push("a")
	s.Push("b")
	s.Push("d")
	s.Push("e")
	s.Push("f")
	s.Print()
	s.Pop()
	s.Pop()
	s.Pop()
	fmt.Println("Stack elements after removeing 3 elements")
	s.Print()
}

func TestPeekElementInStack(t *testing.T) {
	s := newStack()
	s.Push("a")
	s.Push("b")
	s.Push("d")
	s.Push("e")
	s.Push("f")
	s.Print()
	fmt.Println("Stack peek element", s.Peek())
}
