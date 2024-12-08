package stack

import (
	"algorithms/list"
	"errors"
)

// Stack of elements
type Stack[T any] struct {
	list *list.LinkedList[T]
}

// NewStack create an empty stack
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{list: list.NewLinkedList[T]()}
}

// Push a value onto the stack
func (stack *Stack[T]) Push(value T) {
	stack.list.AddFront(value)
}

// Pop a value from the stack, returning an error if the stack is empty
func (stack *Stack[T]) Pop() (*T, error) {
	elem, err := stack.list.RemoveFront()

	if err != nil {
		return nil, errors.New("stack underflow")
	}

	return elem, nil
}

// Peek the top value of the stack, returning an error if the stack is empty
func (stack *Stack[T]) Peek() (*T, error) {
	elem, err := stack.list.RemoveFront()
	if err != nil {
		return nil, errors.New("stack is empty")
	}
	stack.list.AddFront(*elem)
	return elem, nil
}

// IsEmpty if the stack is empty
func (stack *Stack[T]) IsEmpty() bool {
	return stack.list.IsEmpty()
}

// MaxStack of elements which keeps track of the current max
type MaxStack[T ~int | ~float32 | ~float64] struct {
	underlyingStack *Stack[T]
	max             *Stack[T]
}

// NewMaxStack create an empty max-stack
func NewMaxStack[T ~int | ~float32 | ~float64]() *MaxStack[T] {
	return &MaxStack[T]{underlyingStack: NewStack[T](), max: NewStack[T]()}
}

// Push a value on the max-stack
func (stack *MaxStack[T]) Push(value T) {
	if stack.underlyingStack.IsEmpty() {
		stack.underlyingStack.Push(value)
		stack.max.Push(value)
		return
	}

	maxValue, _ := stack.max.Peek()

	if value >= *maxValue {
		stack.underlyingStack.Push(value)
		stack.max.Push(value)
	} else {
		stack.underlyingStack.Push(value)
	}
}

// Pop a value from the max-stack, returning error if the stack is empty
func (stack *MaxStack[T]) Pop() (*T, error) {
	if stack.underlyingStack.IsEmpty() {
		return nil, errors.New("stack underflow")
	}

	elem, _ := stack.underlyingStack.Pop()
	if maxValue, _ := stack.max.Peek(); *elem == *maxValue {
		_, _ = stack.max.Pop()
	}

	return elem, nil
}

// Max value of the stack, returning error if the stack is empty
// Running time: O(1)
func (stack *MaxStack[T]) Max() (*T, error) {
	if stack.underlyingStack.IsEmpty() {
		return nil, errors.New("stack is empty")
	}

	return stack.max.Peek()
}
