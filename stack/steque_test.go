package stack

import "testing"

func TestSteque(t *testing.T) {
	steque := NewSteque[int]()

	steque.Push(1)
	steque.Push(2)
	steque.Push(3)

	steque.Enqueue(1)
	steque.Enqueue(2)
	steque.Enqueue(3)

	expected := []int{3, 2, 1, 1, 2, 3}

	for i := range expected {
		value, _ := steque.Pop()

		if *value != expected[i] {
			t.Errorf("expected %d, got %d", expected[i], *value)
		}
	}

	if !steque.Empty() {
		t.Error("expected empty steque")
	}
}
