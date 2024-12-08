package stack

import "testing"

func TestMaxStack(t *testing.T) {
	maxStack := NewMaxStack[int]()

	maxStack.Push(1)
	maxStack.Push(2)
	maxStack.Push(3)
	maxStack.Push(2)
	maxStack.Push(4)

	if maxValue, _ := maxStack.Max(); *maxValue != 4 {
		t.Errorf("max stack expected 4, got %d", *maxValue)
	}

	_, _ = maxStack.Pop()

	if maxValue, _ := maxStack.Max(); *maxValue != 3 {
		t.Errorf("max stack expected 3, got %d", *maxValue)
	}

	_, _ = maxStack.Pop()

	if maxValue, _ := maxStack.Max(); *maxValue != 3 {
		t.Errorf("max stack expected 3, got %d", *maxValue)
	}

	_, _ = maxStack.Pop()

	if maxValue, _ := maxStack.Max(); *maxValue != 2 {
		t.Errorf("max stack expected 2, got %d", *maxValue)
	}

	_, _ = maxStack.Pop()

	if maxValue, _ := maxStack.Max(); *maxValue != 1 {
		t.Errorf("max stack expected 1, got %d", *maxValue)
	}
}
