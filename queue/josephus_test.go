package queue

import "testing"

func TestJosephus(t *testing.T) {
	result := Josephus(7, 2)
	expectation := []int{1, 3, 5, 0, 4, 2, 6}

	if !eq(result, expectation) {
		t.Errorf("Expected %v, got %v", expectation, result)
	}
}

func eq(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
