package sorting

import (
	"testing"
)

func Test_BubbleSortSorts(t *testing.T) {
	values := []int{2, 4, 1, 10, 7}
	wanted := []int{1, 2, 4, 7, 10}
	BubbleSort(values)

	if !equal(values, wanted) {
		t.Fatalf("Slice was not sorted. Wanted = %v Got = %v", wanted, values)
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}
