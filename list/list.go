package list

import "errors"

type LinkedList[T any] struct {
	head *LinkedListNode[T]
	tail *LinkedListNode[T]
	size int
}

type LinkedListNode[T any] struct {
	elem T
	next *LinkedListNode[T]
}

// NewLinkedList create a new linked list with elements of type T
func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

// AddFront add an element to the front of the list
// Running time: O(1)
func (list *LinkedList[T]) AddFront(elem T) {
	node := &LinkedListNode[T]{elem: elem, next: list.head}

	if list.head == nil {
		list.tail = node
	}

	list.head = node
	list.size++
}

// AddBack add an element to the back of the list
// Running time: O(1)
func (list *LinkedList[T]) AddBack(elem T) {
	if list.head == nil {
		list.head = &LinkedListNode[T]{elem: elem, next: nil}
		list.tail = list.head
	} else {
		list.tail.next = &LinkedListNode[T]{elem: elem, next: nil}
		list.tail = list.tail.next
	}

	list.size++
}

// RemoveFront remove the first element in the list, returning an error if empty
// Running time: O(1)
func (list *LinkedList[T]) RemoveFront() (*T, error) {
	// Empty list
	if list.head == nil {
		return nil, errors.New("list is empty")
	}

	elem := list.head.elem
	list.size--

	// One element in the list
	if list.head == list.tail {
		list.head = nil
		list.tail = nil

		return &elem, nil
	}

	list.head = list.head.next
	return &elem, nil
}

// RemoveBack remove the last element in the list, returning an error if empty
// Running time: O(N), as list is singly-linked
func (list *LinkedList[T]) RemoveBack() (*T, error) {
	// Empty list
	if list.head == nil {
		return nil, errors.New("list is empty")
	}

	elem := list.tail.elem
	list.size--

	// One element in the list
	if list.head == list.tail {
		list.head = nil
		list.tail = nil

		return &elem, nil
	}

	// Find the element right before tail
	curr := list.head.next
	prev := list.head

	for curr != nil {
		prev = curr
		curr = curr.next
	}

	prev.next = nil
	return &elem, nil
}

// Size of the linked list
// Running time: O(1), as size is maintained internally
func (list *LinkedList[T]) Size() int {
	return list.size
}

// IsEmpty if the list is empty
// Running time: O(1), as size is maintained internally
func (list *LinkedList[T]) IsEmpty() bool {
	return list.size == 0
}
