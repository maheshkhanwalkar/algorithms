package queue

// Creative Problem 37 https://algs4.cs.princeton.edu/13stacks/

// Josephus problem of repeatedly eliminating the M-th person from N starting people until only one person is left.
// Returns a list of people eliminated in the order of elimination, with the final element being the position of the
// person not eliminated.
func Josephus(N int, M int) []int {
	result := make([]int, 0, N)
	queue := NewQueue[int]()

	for i := 0; i < N; i++ {
		queue.Enqueue(i)
	}

	for !queue.IsEmpty() {
		for i := 0; i < M-1; i++ {
			elem, _ := queue.Dequeue()
			queue.Enqueue(*elem)
		}

		eliminated, _ := queue.Dequeue()
		result = append(result, *eliminated)
	}

	return result
}
