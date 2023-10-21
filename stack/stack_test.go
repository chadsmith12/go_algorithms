package stack

import (
	"testing"
)

func TestNewStackIsEmpty(t *testing.T) {
	s := New[int]()

	if s.length != 0 {
		t.Errorf("s.length = %d, want %d", s.length, 0)
	}

	if s.head != nil {
		t.Errorf("s.head = %v, want nil", s.head)
	}
}

func TestPush(t *testing.T) {
	s := New[int]()
	s.Push(10)

	if s.length != 1 {
		t.Errorf("s.length = %d, want 1", s.length)
	}
}

func TestPop(t *testing.T) {
	s := New[int]()
	s.Push(42)
	s.Push(69)
	value, hasValue := s.Pop()

	if s.length != 1 {
		t.Errorf("s.length = %d, want 1", s.length)
	}

	if value != 69 {
		t.Errorf("value = %d, want 69", value)
	}

	if !hasValue {
		t.Errorf("value = %v, want true", hasValue)
	}
}

func TestPeek(t *testing.T) {
	s := New[int]()
	s.Push(42)
	value, hasValue := s.Peek()

	if s.length != 1 {
		t.Errorf("s.length = %d, want 1", s.length)
	}

	if value != 42 {
		t.Errorf("value = %d, want 42", value)
	}

	if !hasValue {
		t.Errorf("value = %v, want true", hasValue)
	}
}

func TestPeekOnEmptyGivesNoValue(t *testing.T) {
	s := New[int]()

	value, hasValue := s.Peek()

	if hasValue {
		t.Errorf("hasValue = %v, want false", hasValue)
	}

	if value != 0 {
		t.Errorf("value = %d, want 0", value)
	}
}
