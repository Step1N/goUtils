package queue

import (
	"fmt"
	"testing"
)

func TestAddElementInQueue(t *testing.T) {
	q := newQueue()
	q.Add("a")
	q.Add("b")
	q.Add("d")
	q.Add("e")
	fmt.Println("\n Queue elements ")
	q.Print()
}

func TestRemoveElementInQueue(t *testing.T) {
	q := newQueue()
	q.Add("a")
	q.Add("b")
	q.Add("d")
	q.Add("e")
	q.Print()
	fmt.Println("\n Queue elements ")
	q.Remove()
	q.Remove()
	q.Remove()
	q.Remove()
	q.Print()
}
