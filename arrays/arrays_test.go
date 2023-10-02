package search

import (
	"testing"
)

func Test_LinearSearchReturnsTrue(t *testing.T) {
	items := [4]int{1, 5, 13, 14}
	toFind := 1

	found := LinearSearch(items[:], toFind)

	if !found {
		t.Fatalf(`"1" was not found inside of %v. Should return true.`, items)
	}
}

func Test_LinearSearchReturnsFalse(t *testing.T) {
	items := [4]int{1, 5, 13, 14}
	toFind := 12

	found := LinearSearch(items[:], toFind)

	if found {
		t.Fatalf(`12 was round inside of %v. Should return false.`, items)
	}
}

func Test_BinarySearchReturnsTrue(t *testing.T) {
	items := [4]int{1, 5, 13, 14}
	toFind := 1

	found := BinarySearch(items[:], toFind)

	if !found {
		t.Fatalf(`"1" was not found inside of %v. Should return true.`, items)
	}
}
