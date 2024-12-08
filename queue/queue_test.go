package queue

import "testing"

func TestEnqueueDequeue(t *testing.T) {
	queue := NewQueue[int]()

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	if size := queue.Size(); size != 3 {
		t.Errorf("Queue size should be 3, but got %d", size)
	}

	for i := 1; i <= 3; i++ {
		elem, err := queue.Dequeue()

		if err != nil {
			t.Error(err)
		}

		if *elem != i {
			t.Errorf("Dequeue: expected %d, got %d", i, *elem)
		}
	}

	_, err := queue.Dequeue()

	if err == nil {
		t.Errorf("Dequeue: expected error for dequeue on empty queue, got none")
	}

	if !queue.IsEmpty() {
		t.Errorf("IsEmpty: expected true, got false")
	}
}
