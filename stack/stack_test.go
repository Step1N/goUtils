package stack

import (
	"fmt"
	"testing"
)

func TestPushElementInStack(t *testing.T) {
	s := newStack()
	s.push("a")
	s.push("b")
	s.push("d")
	s.push("e")
	fmt.Println("Stack elements")
	s.print()
}

func TestPopElementInStack(t *testing.T) {
	s := newStack()
	s.push("a")
	s.push("b")
	s.push("d")
	s.push("e")
	s.push("f")
	s.print()
	s.pop()
	s.pop()
	s.pop()
	fmt.Println("Stack elements after removeing 3 elements")
	s.print()
}

func TestPeekElementInStack(t *testing.T) {
	s := newStack()
	s.push("a")
	s.push("b")
	s.push("d")
	s.push("e")
	s.push("f")
	s.print()
	fmt.Println("Stack peek element", s.peek())
}
