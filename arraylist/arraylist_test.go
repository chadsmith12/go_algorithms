package arraylist

import (
	"testing"
)

func TestNewListWithCapacity(t *testing.T) {
	list := NewList[int](10)

	if list.Len() != 0 {
		t.Errorf("Len() = %d, want 0", list.Len())
	}

	if list.Cap() != 10 {
		t.Errorf("Cap() = %d, want 10", list.Cap())
	}
}

func TestPushUpdatesLength(t *testing.T) {
	list := NewList[int](1)

	list.Push(42)

	if list.Len() != 1 {
		t.Errorf("Len() = %d, want 1", list.Len())
	}
}

func TestPopRemovesItemFromBackOfList(t *testing.T) {
	list := NewList[int](3)
	list.Push(42)
	list.Push(69)
	list.Push(420)

	want := list.Pop()

	if want != 420 {
		t.Errorf("value is %d, want 420", want)
	}

	if list.Len() != 2 {
		t.Errorf("Len() is %d, want 2", list.Len())
	}
}

func TestPopRemovesMultipleItems(t *testing.T) {
	list := NewList[int](0)

	list.Push(69)
	list.Push(42)
	list.Push(420)

	list.Pop()
	want := list.Pop()

	if want != 42 {
		t.Errorf("value is %d, want 42", want)
	}

	if list.Len() != 1 {
		t.Errorf("Len() is %d, want 1", list.Len())
	}

}
