package multiarray

import (
	"testing"
)

func areEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for index := range a {
		if a[index] != b[index] {
			return false
		}
	}

	return true
}

func TestInit2DArray(t *testing.T) {
	rows := 2
	cols := 3

	matrix := New[int](rows, cols)

	if matrix.Rows() != rows {
		t.Errorf("Rows() is %d, want %d", matrix.Rows(), rows)
	}

	if matrix.Cols() != cols {
		t.Errorf("Cols() is %d, want %d", matrix.Cols(), cols)
	}
}

func TestGettingItem(t *testing.T) {
	rows := 3
	cols := 3
	valueToSet := 10

	matrix := New[int](rows, cols)
	matrix.Set(1, 0, valueToSet)
	valueGot := matrix.Get(1, 0)

	if valueGot != valueToSet {
		t.Errorf("Get() value = %d, want %d", valueGot, valueToSet)
	}
}

func TestSettingItem(t *testing.T) {
	rows := 3
	cols := 3
	valueToSet := 10

	matrix := New[int](rows, cols)
	matrix.Set(1, 0, valueToSet)
	valueGot := matrix.data[1][0]

	if valueGot != valueToSet {
		t.Errorf("Set(): value = %d, want %d", valueGot, valueToSet)
	}
}

func TestGetRow(t *testing.T) {
	rows := 3
	cols := 2
	valueToSet := 42
	want := [2]int{42, 42}

	matrix := New[int](rows, cols)
	matrix.FillRow(1, valueToSet)
	got := matrix.Row(1)

	if !areEqual(want[:], got) {
		t.Errorf("Row(): value = %v, want %v", got, want)
	}
}

func TestGetCol(t *testing.T) {
	rows := 3
	cols := 2
	valueToSet := 42
	want := [3]int{42, 42, 42}

	matrix := New[int](rows, cols)
	matrix.FillCol(1, valueToSet)
	got := matrix.Col(1)

	if !areEqual(want[:], got) {
		t.Errorf("Col(): value = %v, want %v", got, want)
	}
}
