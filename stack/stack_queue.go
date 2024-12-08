package stack

import "errors"

// Algorithms 1.3 https://algs4.cs.princeton.edu/13stacks/
// Web Exercises 31

// DoubleStackQueue implements a queue using two stacks
type DoubleStackQueue[T any] struct {
	eq  *Stack[T]
	deq *Stack[T]
}

func NewDoubleStackQueue[T any]() *DoubleStackQueue[T] {
	return &DoubleStackQueue[T]{eq: NewStack[T](), deq: NewStack[T]()}
}

func (queue *DoubleStackQueue[T]) Enqueue(elem T) {
	queue.eq.Push(elem)
}

func (queue *DoubleStackQueue[T]) Dequeue() (*T, error) {
	if queue.deq.IsEmpty() {
		if queue.eq.IsEmpty() {
			return nil, errors.New("empty queue")
		}

		shiftOver(queue.eq, queue.deq)
	}

	return queue.deq.Pop()
}

func (queue *DoubleStackQueue[T]) IsEmpty() bool {
	return queue.eq.IsEmpty() && queue.deq.IsEmpty()
}

func shiftOver[T any](eq *Stack[T], deq *Stack[T]) {
	for !eq.IsEmpty() {
		elem, _ := eq.Pop()
		deq.Push(*elem)
	}
}
