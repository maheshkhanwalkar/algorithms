package sort

func MergeSort[T any](arr []T, comparator func(T, T) int) {
	if len(arr) <= 1 {
		return
	}

	mid := len(arr) / 2

	left := arr[:mid]
	right := arr[mid:]

	MergeSort[T](left, comparator)
	MergeSort[T](right, comparator)

	merge(left, right, arr, comparator)
}

func merge[T any](left []T, right []T, combined []T, comparator func(T, T) int) {
	// Nothing to merge
	if len(left) == 0 || len(right) == 0 {
		return
	}

	result := make([]T, 0, len(left)+len(right))

	i := 0
	j := 0

	for i < len(left) && j < len(right) {
		leftElem := left[i]
		rightElem := right[j]

		if comparator(leftElem, rightElem) > 0 {
			result = append(result, rightElem)
			j++
		} else {
			result = append(result, leftElem)
			i++
		}
	}

	var remPos int
	var remSlice []T

	if i < len(left) {
		remPos = i
		remSlice = left
	} else {
		remPos = j
		remSlice = right
	}

	result = append(result, remSlice[remPos:]...)
	copy(combined, result)
}
