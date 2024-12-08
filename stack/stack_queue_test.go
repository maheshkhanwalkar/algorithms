package stack

import "testing"

func TestDoubleStackQueue(t *testing.T) {
	queue := NewDoubleStackQueue[int]()

	queue.Enqueue(1)
	queue.Enqueue(2)

	elem, _ := queue.Dequeue()

	if *elem != 1 {
		t.Errorf("dequeue: expected 1, got %d", *elem)
	}

	queue.Enqueue(3)
	queue.Enqueue(4)

	results := make([]int, 0, 3)

	for !queue.IsEmpty() {
		elem, _ := queue.Dequeue()
		results = append(results, *elem)
	}

	for i := 2; i < 5; i++ {
		if results[i-2] != i {
			t.Errorf("dequeue: expected %d, got %d", i, results[i-2])
		}
	}
}
