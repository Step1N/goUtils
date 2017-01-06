package queue

import "fmt"

type Queue struct {
	input []interface{}
	count int
}

func NewQueue() *Queue {
	q := newQueue()
	return q
}

func newQueue() *Queue {
	in := make([]interface{}, 0)
	return &Queue{input: in, count: -1}
}

func (q *Queue) Add(element interface{}) {
	if q.count == -1 {
		nq := newQueue()
		nq.input = append(nq.input, element)
		q.count = nq.count + 1
		q.input = nq.input
		return
	}

	q.count = q.count + 1
	q.input = append(q.input, element)
}

func (q *Queue) Remove() interface{} {
	if q.count == -1 {
		return -1
	}

	el := q.input[q.count]
	q.input = q.input[:q.count]
	q.count = q.count - 1
	return el
}

func (q *Queue) Print() {
	if q.count == -1 {
		return
	}
	for i := 0; i < len(q.input); i++ {
		fmt.Printf("%v \t", q.input[i])
	}
}
