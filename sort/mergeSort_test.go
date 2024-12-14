package sort

import (
	"math/rand"
	"testing"
)

func TestMergeSort(t *testing.T) {
	arr := []int{9, 8, 5, 10, 11, 8, 7, 1, 3}
	MergeSort(arr, func(i1 int, i2 int) int {
		return i1 - i2
	})

	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			t.Errorf("merge sort error: arr[i] > arr[i+1]: %d > %d", arr[i], arr[i+1])
		}
	}
}

func BenchmarkMergeSort(b *testing.B) {
	arrSize := 10000000
	arr := make([]int, arrSize)

	b.ResetTimer()

	for range b.N {
		for i := 0; i < arrSize; i++ {
			arr[i] = rand.Int()
		}

		MergeSort(arr, func(i1 int, i2 int) int {
			return i1 - i2
		})
	}
}
