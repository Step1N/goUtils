package queue

import (
	"fmt"
	"testing"
)

func TestAddElementInQueue(t *testing.T) {
	q := newQueue()
	q.add("a")
	q.add("b")
	q.add("d")
	q.add("e")
	fmt.Println("\n Queue elements ")
	q.print()
}

func TestRemoveElementInQueue(t *testing.T) {
	q := newQueue()
	q.add("a")
	q.add("b")
	q.add("d")
	q.add("e")
	q.print()
	fmt.Println("\n Queue elements ")
	q.remove()
	q.remove()
	q.remove()
	q.remove()
	q.print()
}
