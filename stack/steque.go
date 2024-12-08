package stack

import "algorithms/list"

type Steque[T any] struct {
	list *list.LinkedList[T]
}

func NewSteque[T any]() *Steque[T] {
	return &Steque[T]{list: list.NewLinkedList[T]()}
}

func (steque *Steque[T]) Push(value T) {
	steque.list.AddFront(value)
}

func (steque *Steque[T]) Pop() (*T, error) {
	return steque.list.RemoveFront()
}

func (steque *Steque[T]) Enqueue(value T) {
	steque.list.AddBack(value)
}

func (steque *Steque[T]) Empty() bool {
	return steque.list.IsEmpty()
}

func (steque *Steque[T]) Size() int {
	return steque.list.Size()
}
