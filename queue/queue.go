package queue

import "algorithms/list"

// Algorithms 1.3 Bags, Queues, and Stacks: https://algs4.cs.princeton.edu/13stacks/
// General implementation

// Queue of elements
type Queue[T any] struct {
	list *list.LinkedList[T]
}

// NewQueue create a new, empty queue
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{list: list.NewLinkedList[T]()}
}

// Enqueue an element
// Running time: O(1)
func (q *Queue[T]) Enqueue(value T) {
	q.list.AddBack(value)
}

// Dequeue an element
// Running time: O(1)
func (q *Queue[T]) Dequeue() (*T, error) {
	return q.list.RemoveFront()
}

// IsEmpty if the queue is empty
// Running time: O(1)
func (q *Queue[T]) IsEmpty() bool {
	return q.list.IsEmpty()
}

// Size of the queue
// Running time: O(1)
func (q *Queue[T]) Size() int {
	return q.list.Size()
}
