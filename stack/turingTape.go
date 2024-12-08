package stack

import "errors"

// Web Exercises 9 https://algs4.cs.princeton.edu/13stacks/

// TuringTape finite tape
type TuringTape struct {
	left   *Stack[int]
	right  *Stack[int]
	active int
}

// NewTuringTape create a new Turing tape of the given size
func NewTuringTape(size int) *TuringTape {
	right := NewStack[int]()

	for range size - 1 {
		right.Push(0)
	}

	return &TuringTape{left: NewStack[int](), active: 0, right: right}
}

// MoveLeft the tape head, returning error if reached the end of the tape
func (tape *TuringTape) MoveLeft() error {
	if tape.left.IsEmpty() {
		return errors.New("left is empty")
	}

	elem, _ := tape.left.Pop()
	tape.right.Push(tape.active)
	tape.active = *elem
	return nil
}

// MoveRight the tape head, returning error if reached the end of the tape
func (tape *TuringTape) MoveRight() error {
	if tape.right.IsEmpty() {
		return errors.New("no cells remaining to the right")
	}

	elem, _ := tape.right.Pop()
	tape.left.Push(tape.active)
	tape.active = *elem
	return nil
}

// Write the value into the active cell
func (tape *TuringTape) Write(value int) {
	tape.active = value
}

// Look at the value in the active cell
func (tape *TuringTape) Look() int {
	return tape.active
}
