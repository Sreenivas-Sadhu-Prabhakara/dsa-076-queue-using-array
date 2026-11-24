package main

import "fmt"

// Queue is a fixed-capacity FIFO queue over a circular buffer: head and a
// count, with indexes wrapping modulo the capacity so no element ever shifts.
type Queue struct {
	data []int
	head int
	size int
}

// NewQueue returns a queue holding at most capacity values.
func NewQueue(capacity int) *Queue {
	return &Queue{data: make([]int, capacity)}
}

// Enqueue appends v at the back; false means the queue is full.
func (q *Queue) Enqueue(v int) bool {
	if q.size == len(q.data) {
		return false
	}
	q.data[(q.head+q.size)%len(q.data)] = v
	q.size++
	return true
}

// Dequeue removes and returns the front value; false means empty.
func (q *Queue) Dequeue() (int, bool) {
	if q.size == 0 {
		return 0, false
	}
	v := q.data[q.head]
	q.head = (q.head + 1) % len(q.data)
	q.size--
	return v, true
}

// Front returns the front value without removing it.
func (q *Queue) Front() (int, bool) {
	if q.size == 0 {
		return 0, false
	}
	return q.data[q.head], true
}

// Size reports how many values are queued.
func (q *Queue) Size() int { return q.size }

func main() {
	q := NewQueue(3)
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	fmt.Println(q.Enqueue(4)) // full
	v, _ := q.Dequeue()
	fmt.Println(v)
	q.Enqueue(4) // wraps into the freed slot
	for q.Size() > 0 {
		v, _ := q.Dequeue()
		fmt.Print(v, " ")
	}
	fmt.Println()
}
