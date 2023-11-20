package linkedlist

import "testing"

func TestLinkedList(t *testing.T) {
	// default list
	list := NewLinkedList[int]()
	if list.length != 0 {
		t.Errorf("NewLinkedList() length = %d; want 0", list.length)
	}
	if list.head != nil {
		t.Errorf("NewLinkedList() head = %v, want nil", list.head)
	}

	// append items to the list
	list.Append(1)
	list.Append(2)
	if list.Len() != 2 {
		t.Errorf("Len() = %d, want 2", list.Len())
	}
	if list.head.value != 1 {
		t.Errorf("head = %d; want 1", list.head.value)
	}

	// Prepend items to the list
	list.Prepend(0)
	if list.Len() != 3 {
		t.Errorf("Len() = %d; want 3", list.Len())
	}
	if list.head.value != 0 {
		t.Errorf("head = %d; want 0", list.head.value)
	}

	err := list.InsertAt(5, 1)
	if err != nil {
		t.Error("InsertAt(5, 1) returned error; wanted nil")
	}
	if list.Len() != 4 {
		t.Errorf("InsertAt(5, 1) length = %d; want 4", list.Len())
	}
	if list.head.next.value != 5 {
		t.Errorf("InsertAt(5, 1) has %d in the 1st index; wanted 5", list.head.next.value)
	}

	err = list.RemoveAt(1)
	if err != nil {
		t.Error("RemoveAt(1) returned error; wanted nil")
	}
	if list.Len() != 3 {
		t.Errorf("RemoveAt(1) length = %d; want 3", list.Len())
	}
	if list.head.next.value != 1 {
		t.Errorf("RemoveAt(1) has %d in the 1st index; wanted 1", list.head.next.value)
	}
}
